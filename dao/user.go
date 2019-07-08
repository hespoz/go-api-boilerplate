package dao

import (
	"fmt"
	"go-api/config"
	"go-api/entities"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func FindAll() (users []entities.User, err error) {

	session := config.Connect().Session.Copy()
	defer func() {
		fmt.Println("Cerrar session")
		session.Close()
	}()

	c := session.DB("tandem-db").C("users")

	err = c.Insert(&entities.User{Email: "asdfsd",
		Password: "asfd"})
	if err != nil {
		log.Fatal(err)
	}

	var result []entities.User
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func CreateUser(newUser *entities.User) error {
	session := config.Connect().Session.Copy()
	defer func() {
		fmt.Println("Cerrar session")
		session.Close()
	}()

	c := session.DB("tandem-db").C("users")
	newUser.Id = bson.NewObjectId()
	err := c.Insert(&newUser)

	if err != nil {
		return err
	}

	return nil
}
