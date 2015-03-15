package main

import (
	"encoding/json"
	"log/syslog"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoItem struct {
	Id   string
	Text string
}

func apiGet(s *mgo.Session, label string, todo string, mlog *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	list, err := dbQuery(s, label)
	if err != nil {
		mlog.Err("Error: db query went bad: " + err.Error())
		return "Error"
	}

	todos := make([]TodoItem, len(list))
	count := 0
	for i, todo := range list {
		todos[i].Text = todo.Text
		todos[i].Id = todo.Id.Hex()
		count += 1
	}
	jsonOut := struct {
		Label string
		Count int
		Todos []TodoItem
	}{label, count, todos}

	out, err := json.Marshal(jsonOut)
	if err != nil {
		mlog.Err("Error: could not properly construct json")
		return "Error"
	}
	return string(out)
}

func apiRemove(s *mgo.Session, label string, id string, mlog *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	err := dbRemove(s, label, MongoTodo{Id: bson.ObjectIdHex(id), Text: ""})
	if err != nil {
		mlog.Err("Error: db remove went bad: " + err.Error())
		return "Error"
	}
	return "Removing: " + id + " from: " + label
}

func apiAdd(s *mgo.Session, label string, todo string, mlog *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	newTodo := MongoTodo{Id: bson.NewObjectId(), Text: todo}
	err := dbInsert(s, label, newTodo)
	if err != nil {
		mlog.Err("Error: db insert went bad: " + err.Error())
		return "Error"
	}
	return newTodo.Id.Hex()
}

func apiCount(s *mgo.Session, mlog *syslog.Writer) string {
	itemCount, listCount, err := dbCountLists(s)
	var countOut struct {
		ItemCount int
		List      []struct {
			Label string
			Count int
		}
	}

	countOut.ItemCount = listCount
	countOut.List = make([]struct {
		Label string
		Count int
	}, len(itemCount))
	i := 0
	for k, v := range itemCount {
		countOut.List[i].Label = k
		countOut.List[i].Count = v
		i += 1
	}

	out, err := json.Marshal(countOut)
	if err != nil {
		mlog.Err("Error: could not properly construct json")
		return "Error"
	}
	return string(out)
}
