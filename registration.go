package main

import (
	"encoding/json"
	"fmt"

	"github.com/ocalvet/registrationapi/models"

	"github.com/ocalvet/registrationapi/database"
)

func main() {
	db := database.New()

	// Write a fish to the database
	idea := models.Idea{
		ID:          "1",
		Title:       "Sample Idea",
		Description: "A simple idea to test application",
	}

	if err := db.Write("idea", idea.ID, idea); err != nil {
		fmt.Println("Error creating idea")
	}

	dbIdea := models.Idea{}
	if err := db.Read("idea", idea.ID, &dbIdea); err != nil {
		fmt.Println("Error reading idea", err)
	}

	fmt.Println("Found " + dbIdea.Title)
	fmt.Println("")
	// Read all ideas from the database, unmarshaling the response.
	records, err := db.ReadAll("idea")
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
	fmt.Println(ideas)
}
