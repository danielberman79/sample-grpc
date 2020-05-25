package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server   serverInfo
	Database databaseInfo
}
type serverInfo struct {
	Host string
	Port string
}

type databaseInfo struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func (d *databaseInfo) Url() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", d.Username, d.Password, d.Host, d.Port, d.DatabaseName)
}

func readConfig() (config config) {
	_, err := toml.DecodeFile("grpcservice/dev.toml", &config)
	if err != nil {
		log.Fatalf("Unable to decode file: %v\n", err)
	}

	return config
}
