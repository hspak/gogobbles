package main

import (
	"encoding/json"
	"log"
	"log/syslog"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
)

type TodoItem struct {
	Id   string
	Text string
}

type ApiFormat struct {
	Label string
	Todos []TodoItem
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	mainLogger, err := syslog.New(syslog.LOG_ERR, "")
	if err != nil {
		log.Fatal("Error: could not start syslog")
	}

	// Site
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", TempIndex{Text: "hihi"})
	})

	m.Get("/:label", func(params martini.Params, r render.Render) {
		label := params["label"]
		list, err := dbQuery(label)
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
		}

		tmplList, err := make([]TodoItem, len(list))
		for i, todo := range list {
			tmplList[i].Id = todo.Id.Hex()
			tmplList[i].Text = todo.Text
		}
		r.HTML(200, "list", TempList{Label: label, Todos: tmplList})
	})

	// API
	m.Get("/get/:label", func(params martini.Params) string {
		label := params["label"]
		if len(label) == 0 {
			return ""
		}

		list, err := dbQuery(label)
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			return "Error"
		}

		todos := make([]TodoItem, len(list))
		for i, todo := range list {
			todos[i].Text = todo.Text
			todos[i].Id = todo.Id.Hex()
		}
		jsonOut := ApiFormat{Label: label, Todos: todos}
		out, err := json.Marshal(jsonOut)
		if err != nil {
			mainLogger.Err("Error: could not properly construct json")
			return "Error"
		}
		return string(out)
	})

	m.Get("/add/:label/:todo", func(params martini.Params) string {
		label := params["label"]
		todo := params["todo"]
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
	})

	m.Get("/remove/:label/:id", func(params martini.Params) string {
		label := params["label"]
		id := params["id"]
		if len(label) == 0 {
			return ""
		}

		err := dbRemove(label, MongoTodo{Id: bson.ObjectIdHex(id), Text: ""})
		if err != nil {
			mainLogger.Err("Error: db remove went bad: " + err.Error())
			return "Error"
		}
		return "Removing: " + id + " from: " + label
	})

	m.Run()
}
