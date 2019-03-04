package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ocalvet/registrationapi/database"
	"github.com/ocalvet/registrationapi/models"
)

func generateIdeaHandler(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling %s", r.Method)
		switch r.Method {
		case http.MethodGet:
			id := strings.TrimPrefix(r.URL.Path, "/api/ideas/")
			if len(id) > 0 {
				log.Printf("Id %s", id)
				i := db.GetIdea(id)
				fmt.Fprintf(w, "Hi With ID: %s", i.ID)
			} else {
				log.Println("Handling getting all ideas")
				ideas := db.GetIdeas()
				fmt.Println(ideas)
				encodedIdeas, err := json.Marshal(ideas)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(encodedIdeas)
			}
		case http.MethodPost:
			log.Println("Handling creation of an idea")
			b, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			idea := models.Idea{}
			if err = json.Unmarshal(b, &idea); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			db.AddIdea(idea)
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func main() {
	db := database.New()
	fmt.Println(db)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ideas/", generateIdeaHandler(db))
	log.Println("Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
