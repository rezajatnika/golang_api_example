package main

import (
	// Standard library packages
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
)

type (
	// User model structure
	User struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func UserShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Stub an example user resource
	user := User{
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

func UserPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Stub an example user to be populated from the body
	user := User{}

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

func UserDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: Only write status for now
	w.WriteHeader(200)
}

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Get a user resource
	router.GET("/users/:id", UserShow)

	// Post user resource
	router.POST("/users", UserPost)

	// Delete a user resource
	router.DELETE("/users/:id", UserDelete)

	// Run http server
	http.ListenAndServe("127.0.0.1:8080", router)
}
