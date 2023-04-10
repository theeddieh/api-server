package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB interface {
	Open() error
	Close() error
}

type DB struct {
	host     string // localhost
	port     string // 8080
	user     string // postgres
	password string // postgres
	name     string // yummly
	scheme   string // public

	db *sqlx.DB
}

func New() *DB {
	fmt.Println("initializing databse")

	db := &DB{
		host:     "localhost",
		port:     "8080",
		user:     "postgres",
		password: "postgres",
		name:     "yummly",
		scheme:   "public",
	}

	return db

}

func (d *DB) Open() error {

	connectionString := ""

	p, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	d.db = p

	return nil
}

func (d *DB) Close() error {
	return nil
}
