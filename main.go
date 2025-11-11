package main

import (
	"flag"
	"log"

	"github.com/codeliger/tygo/tygo"
)

func main() {
	config := flag.String("config", "tygo.yaml", "config file to load")
	flag.Parse()

	tygoConfig := tygo.ReadConfigFromFilePath(*config)
	t := tygo.New(&tygoConfig)

	err := t.Generate()
	if err != nil {
		log.Fatalf("tygo failed: %v", err)
	}
}
