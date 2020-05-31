package database

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Info provides the required configs for setting up a connection to Postgres
type Info struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func (d *Info) url() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", d.Username, d.Password, d.Host, d.Port, d.DatabaseName)
}

//Database represents a wrapper around pgx that has migration abilities
type Database struct {
	*pgxpool.Pool
	url string
}

//Migrate performs Database migrations
func (d *Database) Migrate() error {
	_, b, _, _ := runtime.Caller(0)
	path := fmt.Sprintf("file:///%v/migrations", filepath.Dir(b))

	m, err := migrate.New(path, d.url)

	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		return err
	}
	return nil
}

//Reset drops tables and recreates them
func (d *Database) Reset() error {
	_, b, _, _ := runtime.Caller(0)
	path := fmt.Sprintf("file:///%v/migrations", filepath.Dir(b))

	m, err := migrate.New(path, d.url)
	if err != nil {
		return err
	}

	if err = m.Drop(); err != nil {
		return err
	}

	return d.Migrate()
}

//FromConfig creates a pgx connection from the configuration
func FromConfig(info Info) (*Database, error) {
	c, err := pgxpool.Connect(context.Background(), info.url())

	if err != nil {
		return nil, err
	}

	return &Database{Pool: c, url: info.url()}, nil
}
