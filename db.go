package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoTodo struct {
	Id   bson.ObjectId `bson:"_id"`
	Text string        `bson:"text"`
}

func dbInsert(label string, todo MongoTodo) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("gotest").C(label)
	if err = c.Insert(&todo); err != nil {
		log.Fatal(err)
	}
}

func dbRemove(label string, todo MongoTodo) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("gotest").C(label)
	if err = c.Remove(bson.M{"_id": todo.Id}); err != nil {
		return err
	}
	return nil
}

func dbQuery(label string) []MongoTodo {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	results := make([]MongoTodo, 0)
	c := session.DB("gotest").C(label)
	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Fatal(err)
	}
	return results
}
