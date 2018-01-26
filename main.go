package main

import (
	"os"

	"github.com/fatih/color"
	mgo "gopkg.in/mgo.v2"
)

const mongoip = "127.0.0.1:27017"

func main() {
	session, err := getSession()
	check(err)

	var databases []string
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i > 0 {
				databases = append(databases, arg)
			}
		}
	}
	for _, database := range databases {
		db := getDatabase(session, database)
		color.Yellow("delete database: " + database)
		err := db.DropDatabase()
		check(err)
	}
}

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + mongoip)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
func getDatabase(session *mgo.Session, database string) *mgo.Database {

	D := session.DB(database)
	return D
}
func getCollection(session *mgo.Session, database string, collection string) *mgo.Collection {

	c := session.DB(database).C(collection)
	return c
}
