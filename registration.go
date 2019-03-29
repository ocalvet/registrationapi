package main

import (
	"log"
	"net"
	"net/http"

	"registrationapi/controllers"
	"registrationapi/database"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// RouterDef interface to define a router
type RouterDef interface {
	GET(string, httprouter.Handle)
	POST(string, httprouter.Handle)
	DELETE(string, httprouter.Handle)
}

func main() {
	db := database.New()
	router := httprouter.New()
	controller := controllers.NewRegistrationController(db)
	router.GET("/api/registrations", controller.HandleGetAll)
	router.GET("/api/registrations/:id", controller.HandleGetOne)
	router.POST("/api/registrations", controller.HandleNewRegistration)
	router.DELETE("/api/registrations/:id", controller.HandleDeleteRegistration)
	log.Println("Listening :8080")
	c := cors.AllowAll()

	server := &http.Server{Handler: c.Handler(router)}
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
