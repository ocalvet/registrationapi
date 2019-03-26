package database

import (
	"encoding/json"
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/ocalvet/registrationapi/models"
)

// Storage reads data from a store
type Storage interface {
	AddRegistration(registration models.Registration) models.Registration
	GetRegistration(string) models.Registration
	GetRegistrations() []models.Registration
	DeleteRegistration(string)
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

// AddRegistration adds a registration to the database
func (db DB) AddRegistration(registration models.Registration) models.Registration {
	id, _ := uuid.NewV4()
	registration.ID = id.String()
	if err := db.db.Write("registration", registration.ID, registration); err != nil {
		fmt.Println("Error creating registration")
		return models.Registration{}
	}
	return registration
}

// GetRegistration Gets a registration by id
func (db DB) GetRegistration(id string) models.Registration {
	dbRegistration := models.Registration{}
	if err := db.db.Read("registration", id, &dbRegistration); err != nil {
		fmt.Println("Error reading registration", err)
	}
	return dbRegistration
}

// DeleteRegistration deletes a resource by id
func (db DB) DeleteRegistration(id string) {
	if err := db.db.Delete("registration", id); err != nil {
		fmt.Println("Error deleting registration", err)
	}
}

// GetRegistrations Gets all the registrations
func (db DB) GetRegistrations() []models.Registration {
	// Read all registrations from the database, unmarshaling the response.
	records, err := db.db.ReadAll("registration")
	if err != nil {
		fmt.Println("Error reading registrations", err)
	}

	registrations := []models.Registration{}
	for _, i := range records {
		registrationFound := models.Registration{}
		if err := json.Unmarshal([]byte(i), &registrationFound); err != nil {
			fmt.Println("Error", err)
		}
		registrations = append(registrations, registrationFound)
	}
	return registrations
}
