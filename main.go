package main

import (
	// Standard library packages
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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

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

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Example of Hello, world!
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	// Get a user resource
	router.GET("/users/:id", UserShow)

	// Run http server
	http.ListenAndServe("127.0.0.1:8080", router)
}
