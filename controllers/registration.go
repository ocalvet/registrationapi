package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"registrationapi/database"
	"registrationapi/models"

	"github.com/julienschmidt/httprouter"
	// "github.com/ocalvet/registrationapi/database"
	// "github.com/ocalvet/registrationapi/models"
)

// RegistrationController handles request for registration
type RegistrationController struct {
	db database.Storage
}

// NewRegistrationController creates a registration controller
func NewRegistrationController(db database.Storage) RegistrationController {
	return RegistrationController{db}
}

// HandleGetAll handles getting all registrations request
func (controller RegistrationController) HandleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Handling getting all registrations")
	enableCors(&w)
	registrations := controller.db.GetRegistrations()
	encodedRegistrations, err := json.Marshal(registrations)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(encodedRegistrations)

}

// HandleGetOne handles getting all registrations request
func (controller RegistrationController) HandleGetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.Printf("handling getting registration with %s", id)
	enableCors(&w)
	if len(id) > 0 {
		i := controller.db.GetRegistration(id)
		encodedRegistration, err := json.Marshal(i)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(encodedRegistration)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("add id to request"))
}

// HandleDeleteRegistration handles deleting a registrations request
func (controller RegistrationController) HandleDeleteRegistration(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.Printf("handling deleting registration with id %s", id)
	enableCors(&w)
	if len(id) > 0 {
		controller.db.DeleteRegistration(id)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "add id to request", 500)
	}
}

// HandleNewRegistration handles adding a registration request
func (controller RegistrationController) HandleNewRegistration(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("handling creation of a registration")
	enableCors(&w)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	registration := models.Registration{}
	if err = json.Unmarshal(b, &registration); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	newRegistration := controller.db.AddRegistration(registration)
	encodedRegistration, err := json.Marshal(newRegistration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(encodedRegistration)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
