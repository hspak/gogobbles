package main

import (
	"encoding/json"
	"log/syslog"

	"gopkg.in/mgo.v2/bson"
)

type TodoItem struct {
	Id   string
	Text string
}

func apiGet(label string, todo string, mainLogger *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	list, err := dbQuery(label)
	if err != nil {
		mainLogger.Err("Error: db query went bad: " + err.Error())
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
		mainLogger.Err("Error: could not properly construct json")
		return "Error"
	}
	return string(out)
}

func apiRemove(label string, id string, mainLogger *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	err := dbRemove(label, MongoTodo{Id: bson.ObjectIdHex(id), Text: ""})
	if err != nil {
		mainLogger.Err("Error: db remove went bad: " + err.Error())
		return "Error"
	}
	return "Removing: " + id + " from: " + label
}

func apiAdd(label string, todo string, mainLogger *syslog.Writer) string {
	if len(label) == 0 {
		return ""
	}

	newTodo := MongoTodo{Id: bson.NewObjectId(), Text: todo}
	err := dbInsert(label, newTodo)
	if err != nil {
		mainLogger.Err("Error: db insert went bad: " + err.Error())
		return "Error"
	}
	return newTodo.Id.Hex()
}

func apiCount(mainLogger *syslog.Writer) string {
	itemCount, listCount, err := dbCountLists()
	var countOut struct {
		ListCount int
		List      []struct {
			Label string
			Count int
		}
	}

	countOut.ListCount = listCount
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
		mainLogger.Err("Error: could not properly construct json")
		return "Error"
	}
	return string(out)
}
