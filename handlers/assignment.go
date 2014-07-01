package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignment ...
func CreateAssignment(c *gin.Context) {
	a := new(m.Assignment)

	errs := binding.Bind(c.Req, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Assignments.Store(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = id

	writeJSON(c.Writer, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAssignment ...
func GetAssignment(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.Assignment{}
	err := store.Assignments.FindByID(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	writeJSON(c.Writer, &APIRes{"assignment": []m.Assignment{a}})
	return
}

// UpdateAssignment ...
func UpdateAssignment(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.Assignment)

	errs := binding.Bind(c.Req, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	a.ID = id
	err := store.Assignments.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating Assignment", 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAllAssignments ...
func GetAllAssignments(c *gin.Context) {
	assignment := []m.Assignment{}
	err := store.Classes.FindAll(&assignment)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"assignment": assignment})
	return
}
