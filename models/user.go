package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User model structure
	User struct {
		Id   bson.ObjectId `json:"id"   bson:"_id"`
		Name string        `json:"name" bson:"name"`
		Age  int           `json:"age"  bson:"age"`
	}

	// User collection
	UserRepo struct {
		collection *mgo.Collection
	}
)

func getCollection() *mgo.Collection {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s.DB("go_api_example").C("users")
}

func NewUserRepo() *UserRepo {
	return &UserRepo{getCollection()}
}

func (ur *UserRepo) Find(id string) User {
	res := User{}
	oid := bson.ObjectIdHex(id)
	ur.collection.FindId(oid).One(&res)
	return res
}

func (ur *UserRepo) Ind() []User {
	res := []User{}
	ur.collection.Find(nil).All(&res);
	return res
}
