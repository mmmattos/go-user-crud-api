package main

import "gopkg.in/mgo.v2/bson"

// User Model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Age      string        `bson:"age" json:"age"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Address  string        `bson:"address" json:"address"`
}
