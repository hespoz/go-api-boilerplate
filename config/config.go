package config

import (
	"log"
	"os"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var once sync.Once

type SessionDB struct {
	Session *mgo.Session
}

var sessionDB *SessionDB

func Connect() *SessionDB {

	once.Do(func() {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{os.Getenv("MONGO_DB")},
			Timeout:  60 * time.Second,
			Database: os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("PASSWORD"),
		}

		log.Print("Creating session Object")
		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		sessionDB = &SessionDB{Session: session}
		if err != nil {
			panic(err)
		}
	})

	return sessionDB
}
