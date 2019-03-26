package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ocalvet/registrationapi/controllers"
	"github.com/ocalvet/registrationapi/database"
)

// RouterDef interface to define a router
type RouterDef interface {
	GET(string, httprouter.Handle)
	POST(string, httprouter.Handle)
	DELETE(string, httprouter.Handle)
}

func generateIdeaHandler(router RouterDef, db database.DB) {
	controller := controllers.NewRegistrationController(db)
	router.GET("/api/registrations", controller.HandleGetAll)
	router.GET("/api/registrations/:id", controller.HandleGetOne)
	router.POST("/api/registrations", controller.HandleNewRegistration)
	router.DELETE("/api/registrations/:id", controller.HandleDeleteRegistration)
}

func main() {
	db := database.New()
	router := httprouter.New()
	generateIdeaHandler(router, db)
	log.Println("Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
