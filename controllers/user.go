package controllers

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
	"github.com/rezajatnika/golang_api_example/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (uc UserController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Grab user id
	id := ps.ByName("id")

	// Verify id is ObjectId otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab user object id
	objectId := bson.ObjectIdHex(id)

	// Stub user first
	user := models.User{}

	// Fetch user
	err := uc.session.DB("go_api_example").C("users").FindId(objectId).One(&user)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not Found\n")
		return
	}

	// Return user JSON
	uj, _ := json.Marshal(user)

	// Render as JSON with header
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Stub an example user to be populated from the body
	user := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&user)

	// Add user id
	user.Id = bson.NewObjectId()

	// Return user JSON encoding
	uj, _ := json.Marshal(user)

	// Write user to MongoDB
	uc.session.DB("go_api_example").C("users").Insert(user)

	// Render as JSON with header
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Grab user id
	id := ps.ByName("id")

	// Verify id is ObjectId otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab object id
	objectId := bson.ObjectIdHex(id)

	// Fetch user
	err := uc.session.DB("go_api_example").C("users").RemoveId(objectId)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not Found\n")
		return
	}

	// Render as JSON with header
	w.WriteHeader(200)
}
