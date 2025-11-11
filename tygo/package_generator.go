package tygo

import (
	"go/ast"
	"go/token"
	"os"
	"strings"
)

// writeGeneratedFile writes the generated code for a single file to the given strings.Builder.
func (g *PackageGenerator) writeGeneratedFile(s *strings.Builder, file *ast.File, outFilePath string) {
	first := true

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		// GenDecl can be an import, type, var, or const expression
		case *ast.GenDecl:
			if x.Tok == token.IMPORT {
				return false
			}
			isEmit := false
			if x.Tok == token.VAR {
				isEmit = g.isEmitVar(x)
				if !isEmit {
					return false
				}
			}

			if first {
				if outFilePath != "" {
					g.writeFileSourceHeader(s, outFilePath, file)
				}
				first = false
			}
			if isEmit {
				g.emitVar(s, x)
				return false
			}
			g.writeGroupDecl(s, x)
			return false
		}
		return true
	})
}

func (g *PackageGenerator) Generate() (string, error) {
	s := new(strings.Builder)

	g.writeFileCodegenHeader(s)
	g.writeFileFrontmatter(s)

	filepaths := g.GoFiles

	for _, filepath := range filepaths {
		if g.conf.IsFileIgnored(filepath) {
			continue
		}

		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", err
		}

		ts, err := GenerateFromString(string(content), g.conf)
		if err != nil {
			return "", err
		}

		s.WriteString(ts)
	}

	return s.String(), nil
}
