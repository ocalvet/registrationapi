package database

import (
	"encoding/json"
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/ocalvet/registrationapi/models"
)

// StoreReader reads data from a store
type StoreReader interface {
	Get(id string) interface{}
}

// DB is the database
type DB struct {
	db *scribble.Driver
}

// New creates a new database
func New() DB {
	dir := "./data"
	db := DB{}
	scribble, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Printf("Error creating database file %v", err)
	}
	db.db = scribble
	return db
}

func (db DB) AddIdea(idea models.Idea) {
	if err := db.db.Write("idea", idea.ID, idea); err != nil {
		fmt.Println("Error creating idea")
	}
}

func (db DB) GetIdea(id string) models.Idea {
	dbIdea := models.Idea{}
	if err := db.db.Read("idea", id, &dbIdea); err != nil {
		fmt.Println("Error reading idea", err)
	}
	return dbIdea
}

func (db DB) GetIdeas() []models.Idea {
	// Read all ideas from the database, unmarshaling the response.
	records, err := db.db.ReadAll("idea")
	if err != nil {
		fmt.Println("Error reading ideas", err)
	}

	ideas := []models.Idea{}
	for _, i := range records {
		ideaFound := models.Idea{}
		if err := json.Unmarshal([]byte(i), &ideaFound); err != nil {
			fmt.Println("Error", err)
		}
		ideas = append(ideas, ideaFound)
	}
	return ideas
}
