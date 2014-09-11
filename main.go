package main

import (
	"encoding/json"

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

	// Site
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", TempIndex{Text: "hihi"})
	})

	m.Get("/:label", func(params martini.Params, r render.Render) {
		label := params["label"]
		list := dbQuery(label)
		tmplList := make([]TodoItem, len(list))
		for i, todo := range list {
			tmplList[i].Id = todo.Id.Hex()
			tmplList[i].Text = todo.Text
		}
		r.HTML(200, "list", TempList{Label: label, Todos: tmplList})
	})

	// API
	m.Get("/get/:label", func(params martini.Params) string {
		label := params["label"]
		list := dbQuery(label)
		todos := make([]TodoItem, len(list))
		for i, todo := range list {
			todos[i].Text = todo.Text
			todos[i].Id = todo.Id.Hex()
		}
		jsonOut := ApiFormat{Label: label, Todos: todos}
		out, _ := json.Marshal(jsonOut)
		return string(out)
	})

	m.Get("/add/:label/:todo", func(params martini.Params) string {
		label := params["label"]
		todo := params["todo"]
		newTodo := MongoTodo{Id: bson.NewObjectId(), Text: todo}
		dbInsert(label, newTodo)
		return newTodo.Id.Hex()
	})

	m.Get("/remove/:label/:id", func(params martini.Params) string {
		label := params["label"]
		id := params["id"]
		err := dbRemove(label, MongoTodo{Id: bson.ObjectIdHex(id), Text: ""})
		if err != nil {
			return id + " does not exist in " + label
		}
		return "Removing: " + id + " from: " + label
	})

	m.Run()
}
