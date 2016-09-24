package controllers

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
	"github.com/rezajatnika/golang_api_example/models"
)

type (
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) Show(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	// Stub an example user resource
	user := models.User{
		Id:   ps.ByName("id"),
		Name: "John Doe",
		Age:  20,
	}

	// Return user JSON encoding
	uj, _ := json.Marshal(user)

	// Render as JSON with header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Stub an example user to be populated from the body
	user := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&user)

	// Add user id
	user.Id = "1234"

	// Return user JSON encoding
	uj, _ := json.Marshal(user)

	// Render as JSON with header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: Only write status for now
	w.WriteHeader(200)
}
