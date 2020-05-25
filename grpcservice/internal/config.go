package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/djquan/skeleton/grpcservice/internal/platform/database"
	"log"
)

//Config represents the configuration of this application
type Config struct {
	Server   serverInfo
	Database database.Info
}

type serverInfo struct {
	Host string
	Port string
}

//ReadConfig parses a config file and returns a representation
func ReadConfig() (config Config) {
	_, err := toml.DecodeFile("grpcservice/dev.toml", &config)
	if err != nil {
		log.Fatalf("Unable to decode file: %v\n", err)
	}

	return config
}
