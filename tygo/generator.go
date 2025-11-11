package tygo

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
	"gopkg.in/yaml.v3"
)

// Generator for one or more input packages, responsible for linking
// them together if necessary.
type Tygo struct {
	conf *Config

	packageGenerators map[string]*PackageGenerator
}

// Responsible for generating the code for an input package
type PackageGenerator struct {
	conf    *PackageConfig
	pkg     *packages.Package
	GoFiles []string
}

func New(config *Config) *Tygo {
	return &Tygo{
		conf:              config,
		packageGenerators: make(map[string]*PackageGenerator),
	}
}

func (g *Tygo) SetTypeMapping(goType string, tsType string) {
	for _, p := range g.conf.Packages {
		p.TypeMappings[goType] = tsType
	}
}

func (g *Tygo) Generate() error {
	ctx := context.Background()
	pkgs, err := packages.Load(&packages.Config{
		Context: ctx,
		Mode:    packages.NeedSyntax | packages.NeedFiles,
	}, g.conf.PackageNames()...)
	if err != nil {
		return err
	}

	for i, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return fmt.Errorf("%+v", pkg.Errors)
		}

		if len(pkg.GoFiles) == 0 {
			return fmt.Errorf("no input go files for package index %d", i)
		}

		pkgConfig := g.conf.PackageConfig(pkg.ID)

		pkgGen := &PackageGenerator{
			conf:    pkgConfig,
			GoFiles: pkg.GoFiles,
			pkg:     pkg,
		}
		g.packageGenerators[pkg.PkgPath] = pkgGen
		code, err := pkgGen.Generate()
		if err != nil {
			return err
		}

		outPath := pkgGen.conf.ResolvedOutputPath(filepath.Dir(pkg.GoFiles[0]))
		err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(outPath, []byte(code), 0o664)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadConfigFromFilePath(cfgFilePath string) Config {
	data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}

	for _, p := range config.Packages {
		err := p.Normalize()
		if err != nil {
			log.Fatalf("error normalizing config: %v", err)
		}
	}

	return config
}
