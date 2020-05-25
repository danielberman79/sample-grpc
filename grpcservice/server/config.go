package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server serverInfo
}
type serverInfo struct {
	Host string
	Port string
}

func readConfig() (config config) {
	_, err := toml.DecodeFile("grpcservice/dev.toml", &config)
	if err != nil {
		log.Fatalf("Unable to decode file: %v\n", err)
	}

	return config
}
