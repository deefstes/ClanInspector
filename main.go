package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var (
	config       Configuration
	mongoSession *mgo.Session
)

func main() {
	var err error
	config, err = ReadConfig()
	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	// Connect to MongoDB
	mongoSession, err = mgo.Dial(config.MongoDB)
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	defer mongoSession.Close()

	RetrieveMembers()
	RetrieveActivities()

	//FixActivity("1580510807")
}
