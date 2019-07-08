package entities

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `bson:"email" validate:"required"`
	Password string        `bson:"password"`
}
