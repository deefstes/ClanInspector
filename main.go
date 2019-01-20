package main

import (
	"encoding/json"
	"fmt"
	"time"

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

	fmt.Printf("%s | ClanInspector started\r\n", time.Now().Format("2006-01-02 15:04:05"))

	RetrieveMembers()
	//RetrieveActivities()
	//RetrievePlayersStats()
	//RetrievePlayersAggregateStats()

	//FixActivities()

	//WhoPlaysWithWho(
	//	time.Date(2018, time.May, 12, 0, 0, 0, 0, time.UTC),
	//	time.Date(2019, time.June, 12, 0, 0, 0, 0, time.UTC),
	//	time.Now().Format("060102"),
	//	true,
	//)

	//WhoPlaysWhen(
	//	time.Date(2017, time.April,.y.yb,b, ,/, ,h,,/,, , r  dt=0-]5 1, 0, 0, 0, 0, time.UTC),
	//	time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC),
	//	time.Now().Format("060102"),
	//)

	//TestActivity("1676662159")
}

func StructToJSON(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(b)
}
