package main

import (
	"flag"
	"log"

	"github.com/djquan/skeleton/commentservice/internal"
	"github.com/djquan/skeleton/commentservice/internal/platform/database"
)

func main() {
	command := flag.String("command", "", "migrate/reset")
	flag.Parse()

	config := internal.ReadConfig()

	switch *command {
	case "migrate":
		db, err := database.FromConfig(config.Database)
		if err != nil {
			log.Fatalf("Unable to talk to database: %v", err)
		}

		log.Println("Performing Database Migration")

		if err = db.Migrate(); err != nil {
			log.Fatalf("Unable to migrate database: %v", err)
		}
	case "reset":
		db, err := database.FromConfig(config.Database)
		if err != nil {
			log.Fatalf("Unable to talk to database: %v", err)
		}
		log.Println("Performing Database Reset")

		if err = db.Reset(); err != nil {
			log.Fatalf("Unable to reset database: %v", err)
		}

	default:
		log.Fatalf("Provide an appropriate --command")
	}
}
