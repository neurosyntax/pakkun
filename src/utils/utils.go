/*
    utils.go

    Utilities for handling data
    
    Author: Justin Chen
    2.14.2017

    Boston University 
    Computer Science

    Dependencies:        exuberant ctags, and mongodb driver for go (http://labix.org/mgo)
    Operating systems:   GNU Linux, OS X
    Supported languages: C, C++, C#, Erlang, Lisp, Lua, Java, Javascript, and Python
*/

package utils

import (
	"log"
	"gopkg.in/mgo.v2"
)

/*
	Connect to MongoDB and return the session
	User needs to handle Session.Close()
*/
func ConnectDB() *mgo.Session {
	// Connect to MongoDB
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    return session
}

func SaveMgoDoc(dbName string, collectionName string, file interface{}) bool {
    session, err := mgo.Dial("localhost:27017")
    
    if err != nil {
        panic(err)
    }
    
    defer session.Close()

    if err != nil {
        log.Printf("failed to marshal struct to json...\n", file)
        return false
    }

    collection := session.DB(dbName).C(collectionName)
    err         = collection.Insert(&file)

    if err != nil {
        log.Printf("failed to insert doc into database...\n", file)
        return false
    }

    return true
}