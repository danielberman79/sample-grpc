package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/djquan/skeleton/grpcservice/internal/platform/database"
	"github.com/google/uuid"
)

//NewDatabaseForTest provides a new database for tests, and returns that database and the cleanup function.
func NewDatabaseForTest(t *testing.T) (database.Database, func()) {
	config := database.Info{
		Host:         "localhost",
		Port:         "5432",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "postgres",
	}

	db, err := database.FromConfig(config)

	if err != nil {
		t.Fatalf("Unable to establish connection to database in order to create test database: %v\n", err)
	}

	tableId := strings.ReplaceAll(fmt.Sprintf("skeleton-test-%v", uuid.New().String()), "-", "")
	_, err = db.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %v", tableId))

	if err != nil {
		cleanup(nil, db, t, tableId)()
		t.Fatalf("Unable to create test database, %v, %v\n", tableId, err)
	}

	config = database.Info{
		Host:         "localhost",
		Port:         "5432",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: tableId,
	}

	newDb, err := database.FromConfig(config)

	if err != nil {
		cleanup(nil, db, t, tableId)()
		t.Fatalf("Unable to establish connection with new test database, %v", err)
	}

	if err = newDb.Migrate(); err != nil {
		cleanup(newDb, db, t, tableId)()
		t.Fatalf("Unable to perform DB migrations: %v\n", err)
	}

	return *newDb, cleanup(newDb, db, t, tableId)
}

func cleanup(newDb, db *database.Database, t *testing.T, tableId string) func() {
	return func() {
		if newDb != nil {
			newDb.Close()
		}

		_, err := db.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %v", tableId))
		if err != nil {
			t.Fatalf("Unable to drop test database %v: %v", tableId, err)
		}

		db.Close()
	}
}
