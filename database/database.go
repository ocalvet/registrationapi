package database

import (
	"fmt"

	scribble "src/github.com/nanobox-io/golang-scribble"
)

// StoreReader reads data from a store
type StoreReader interface {
	Get(id string) interface{}
}

// DB is the database
type DB struct {
	db scribble.DB
}

// New creates a new database
func New() DB {
	dir := "./data"

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Fatal("Error creating database file", err)
	}

	return DB{}
}
