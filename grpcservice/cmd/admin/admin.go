package main

import (
	"flag"
	"github.com/djquan/skeleton/grpcservice/internal"
	"github.com/djquan/skeleton/grpcservice/internal/platform/database"
	"log"
)

func main() {
	command := flag.String("command", "", "migrate/")
	flag.Parse()

	config := internal.ReadConfig()

	switch *command {
	case "migrate":
		db, err := database.FromConfig(config.Database)
		if err != nil {
			log.Fatalf("Unable to talk to database: %v", err)
		}

		log.Println("Performing Database Migration")
		err = db.Migrate()

		if err != nil {
			log.Fatalf("Unable to migrate database: %v", err)
		}
	default:
		log.Fatalf("Provide an appropriate --command")
	}
}
