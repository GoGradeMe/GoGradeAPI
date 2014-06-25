package handlers

import (
	d "github.com/Lanciv/GoGradeAPI/database"
	// "github.com/Lanciv/GoGradeAPI/model"

	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := d.GetAllUsers()
	if err != nil {
		writeError(w, serverError, 500)
		return
	}
	writeJSON(w, users)

}
