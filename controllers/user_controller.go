package controllers

import (
	"go-api/dao"
	"go-api/entities"
	"log"
)

func FindAll() (users []entities.User, err error) {

	result, err := dao.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func CreateUser(user *entities.User) error {
	return dao.CreateUser(user)
}
