package main

import (
	"fmt"

	"github.com/ocalvet/registrationapi/database"
	"github.com/ocalvet/registrationapi/models"
)

func main() {
	db := database.New()
	fmt.Println(db)

	// Add an idea to the database
	idea := models.Idea{
		ID:          "1",
		Title:       "Sample Idea",
		Description: "A simple idea to test application",
	}

	db.AddIdea(idea)

	i := db.GetIdea(idea.ID)
	fmt.Println("Found " + i.Title)

	ideas := db.GetIdeas()
	fmt.Println(ideas)
}
