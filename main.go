package main

import (
	// Standard library packages
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
	"github.com/rezajatnika/golang_api_example/controllers"
	"github.com/rezajatnika/golang_api_example/lib"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial(lib.Mong)
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// User controller
	userController := controllers.NewUserController(getSession())

	// User resource
	router.GET("/users/:id", userController.Show)
	router.POST("/users", userController.Create)
	router.PATCH("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)

	// API
	router.GET("/v1/users/:id", userController.Read)
	router.GET("/v1/users", userController.Index)

	// Run http server
	http.ListenAndServe(":" + lib.Port, router)
}
