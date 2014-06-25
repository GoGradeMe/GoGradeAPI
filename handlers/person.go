package handlers

import (
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/gorilla/mux"
	"net/http"
)

// CreatePerson allows you to create a Person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p m.Person

	if readJSON(r, &p) {
		// Person should exist before trying to do anything.
		// if (p == m.Person{}) {
		// 	writeError(w, "Person required", 400)
		// 	return
		// }

		err := d.CreatePerson(&p)
		if err != nil {
			writeError(w, "Error creating Person", 500)
			return
		}

	} else {
		writeError(w, "Error parsing JSON", 400)
		return
	}

	writeJSON(w, p)
	return
}

// GetPerson will return a Person with all of their Profiles.
func GetPerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pID, ok := vars["id"]
	if !ok {
		writeError(w, "Invalid Person ID", 400)
		return
	}

	p, err := d.GetPerson(pID)
	if err != nil {
		writeError(w, serverError, 400)
		return
	}

	writeJSON(w, p)
	return
}

// GetAllPeople returns all people without their profiles.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := d.GetAllPeople()
	if err != nil {
		writeError(w, serverError, 500)
		return
	}
	writeJSON(w, people)
	return
}
