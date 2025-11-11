package config

import (
	"log"
	"os"

	"github.com/codeliger/tygo/tygo"
	"gopkg.in/yaml.v3"
)

func ReadFromFilepath(cfgFilepath string) tygo.Config {
	b, err := os.ReadFile(cfgFilepath)
	if err != nil {
		log.Fatalf("could not read config file from %s: %v", cfgFilepath, err)
	}
	conf := tygo.Config{}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("could not parse config file from: %v", err)
	}

	return conf
}
