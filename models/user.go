package models

import "gopkg.in/mgo.v2/bson"

type (
	// User model structure
	User struct {
		Id   bson.ObjectId `json:"id"   bson:"_id"`
		Name string        `json:"name" bson:"name"`
		Age  int           `json:"age"  bson:"age"`
	}
)
