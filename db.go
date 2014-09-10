package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Text string
}

func dbInsert(label string, todo Todo) {
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

func dbRemove(label string, todo Todo) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("gotest").C(label)
	if err = c.Remove(bson.M{"text": todo.Text}); err != nil {
		return err
	}
	return nil
}

func dbQuery(label string) []Todo {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	results := make([]Todo, 0)
	c := session.DB("gotest").C(label)
	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Fatal(err)
	}
	return results
}
