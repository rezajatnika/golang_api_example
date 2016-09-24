package main

import (
	// Standard library packages
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
	"github.com/rezajatnika/golang_api_example/controllers"
)

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// User controller
	userController := controllers.NewUserController()

	// User resource
	router.GET("/users/:id", userController.Show)
	router.POST("/users", userController.Create)
	router.DELETE("/users/:id", userController.Delete)

	// Run http server
	http.ListenAndServe("127.0.0.1:8080", router)
}
