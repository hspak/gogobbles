package main

import (
	"log"
	"log/syslog"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// ~150k list limit

// this is a monster
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

	m.Get("/list/:label", func(params martini.Params, r render.Render) {
		tmplList, err := getListValues(params["label"], mainLogger)
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			r.HTML(500, "index", nil)
		}
		r.HTML(200, "list", TempList{Label: params["label"], Todos: tmplList})
	})

	// API
	m.Get("/api/get/:label", func(params martini.Params) string {
		return apiGet(params["label"], params["todo"], mainLogger)
	})

	m.Get("/api/remove/:label/:id", func(params martini.Params) string {
		return apiRemove(params["label"], params["id"], mainLogger)
	})

	m.Get("/api/add/:label/:todo", func(params martini.Params) string {
		return apiAdd(params["label"], params["todo"], mainLogger)
	})

	m.Get("/api/count", func() string {
		return apiCount(mainLogger)
	})

	m.Run()
}
