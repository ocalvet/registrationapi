package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ocalvet/registrationapi/database"
)

func generateIdeaHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := strings.TrimPrefix(r.URL.Path, "/api/ideas/")
			if len(id) > 0 {
				log.Printf("Id %s", id)
				fmt.Fprintf(w, "Hi With ID: %s", id)
			} else {
				log.Println("Handling GET")
				fmt.Fprint(w, "Hi Without ID")
			}
		case http.MethodPost:
			log.Println("Handling POST")
			fmt.Fprint(w, "Hi Post")
		default:
			w.WriteHeader(404)
			fmt.Fprintf(w, "Wrong method: %s", r.Method)
		}
	}
}

func main() {
	db := database.New()
	fmt.Println(db)

	// // Add an idea to the database
	// idea := models.Idea{
	// 	ID:          "1",
	// 	Title:       "Sample Idea",
	// 	Description: "A simple idea to test application",
	// }

	// db.AddIdea(idea)

	// i := db.GetIdea(idea.ID)
	// fmt.Println("Found " + i.Title)

	// ideas := db.GetIdeas()
	// fmt.Println(ideas)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ideas/", generateIdeaHandler())
	log.Println("Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
