package main

import (
	"encoding/json"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type OutList struct {
	Label string
	Todos []string
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "hi?")
	})

	m.Get("/get/:label", func(params martini.Params) string {
		list := dbQuery(params["label"])
		todos := make([]string, 0)
		for _, todo := range list {
			todos = append(todos, todo.Text)
		}
		outlist := OutList{Label: params["label"], Todos: todos}
		out, _ := json.Marshal(outlist)
		return string(out)
	})

	m.Get("/add/:label/:todo", func(params martini.Params) string {
		dbInsert(params["label"], Todo{Text: params["todo"]})
		return "Adding: " + params["label"]
	})

	m.Get("/remove/:label/:todo", func(params martini.Params) string {
		err := dbRemove(params["label"], Todo{Text: params["todo"]})
		if err != nil {
			return
		}
		return "Removing: " + params["todo"] + " from: " + params["label"]
	})

	m.Get("/:label", func(params martini.Params) string {
		list := dbQuery(params["label"])
		str := params["label"] + " items:\n"
		for _, todo := range list {
			str += "    " + todo.Text
		}
		return str
	})

	m.Run()
}
