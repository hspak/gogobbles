package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoTodo struct {
	Id   bson.ObjectId `bson:"_id"`
	Text string        `bson:"text"`
}

func dbInsert(label string, todo MongoTodo) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("gotest").C(label)
	if err = c.Insert(&todo); err != nil {
		return err
	}
	return nil
}

func dbRemove(label string, todo MongoTodo) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("gotest").C(label)
	if err = c.Remove(bson.M{"_id": todo.Id}); err != nil {
		return err
	}
	return nil
}

func dbQuery(label string) ([]MongoTodo, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	results := make([]MongoTodo, 0)
	c := session.DB("gotest").C(label)
	if err = c.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}
