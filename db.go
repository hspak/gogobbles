package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoTodo struct {
	Id   bson.ObjectId `bson:"_id"`
	Text string        `bson:"text"`
}

func dbOpen() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func dbInsert(session *mgo.Session, label string, todo MongoTodo) error {
	c := session.DB("gotest").C(label)
	if err := c.Insert(&todo); err != nil {
		return err
	}
	return nil
}

func dbRemove(session *mgo.Session, label string, todo MongoTodo) error {
	c := session.DB("gotest").C(label)
	if err := c.Remove(bson.M{"_id": todo.Id}); err != nil {
		return err
	}
	return nil
}

func dbQuery(session *mgo.Session, label string) ([]MongoTodo, error) {
	results := make([]MongoTodo, 0)
	c := session.DB("gotest").C(label)
	if err := c.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func dbCountLists(session *mgo.Session) (map[string]int, int, error) {
	listNames, err := session.DB("gotest").CollectionNames()
	if err != nil {
		return nil, 0, err
	}

	itemCount := make(map[string]int)
	var listCount int
	for _, name := range listNames {
		if name == "system.indexes" {
			continue
		}

		count, err := session.DB("gotest").C(name).Count()
		if err != nil {
			return nil, 0, err
		}
		itemCount[name] = count
		listCount += count
	}
	return itemCount, listCount, nil
}
