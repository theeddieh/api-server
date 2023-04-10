package database

import (
	"api-server/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB interface {
	Open() error
	Close() error
	GetUsers() ([]*models.User, error)
}

type DB struct {
	host       string // localhost
	port       int    // 8080
	user       string // postgres
	password   string // postgres
	name       string // yummly
	scheme     string // public
	connection string

	db *sqlx.DB
}

func New() *DB {
	fmt.Println("initializing database")

	db := &DB{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		name:     "yummly",
		scheme:   "public",
	}

	c := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host,
		db.port,
		db.user,
		db.password,
		db.name)

	db.connection = c

	return db

}

func (d *DB) Open() error {

	fmt.Println("opening connection to postgres")
	fmt.Println(d.connection)

	p, err := sqlx.Open("postgres", d.connection)

	if err != nil {
		return err
	}

	d.db = p

	fmt.Println("pinging postgres")
	err = d.db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("connected to postgres")

	return nil
}

func (d *DB) Close() error {
	fmt.Println("closing postgres connection")
	return d.db.Close()
}

func (d *DB) GetUsers() ([]*models.User, error) {
	var chefs []*models.User

	query := "SELECT * FROM users"

	err := d.db.Select(&chefs, query)

	return chefs, err
}
