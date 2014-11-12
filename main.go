package main

import (
	"log"
	"log/syslog"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// ~150k list limit

func maxLen(s string) int {
	if len(s) > 80 {
		return 80
	}
	return len(s)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	if err := dbCheck(); err != nil {
		log.Fatal("Error: could not connect to mongodb")
	}

	mainLogger, err := syslog.New(syslog.LOG_ERR, "")
	if err != nil {
		log.Fatal("Error: could not start syslog")
	}

	// Site
	m.Get("/", func(r render.Render) {
		count, err := getIndexInfo()
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			r.HTML(500, "index", nil) // make a new tmpl for this
		}
		r.HTML(200, "index", TempIndex{ListCount: strconv.Itoa(count)})
	})

	m.Get("/api", func(r render.Render) {
		r.HTML(200, "api", nil)
	})
	m.Get("/faq", func(r render.Render) {
		r.HTML(200, "faq", nil)
	})

	m.Get("/list/:label", func(params martini.Params, r render.Render) {
		label := params["label"][:maxLen(params["label"])]
		tmplList, err := getListValues(label)
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			r.HTML(500, "index", nil) // make a new tmpl for this
		}
		r.HTML(200, "list", TempList{Label: label, Todos: tmplList})
	})

	// API
	m.Get("/api/get/:label", func(params martini.Params) string {
		label := params["label"][:maxLen(params["label"])]
		todo := params["todo"][:maxLen(params["todo"])]
		return apiGet(label, todo, mainLogger)
	})

	m.Get("/api/remove/:label/:id", func(params martini.Params) string {
		label := params["label"][:maxLen(params["label"])]
		id := params["id"][:maxLen(params["id"])]
		return apiRemove(label, id, mainLogger)
	})

	m.Get("/api/add/:label/:todo", func(params martini.Params) string {
		label := params["label"][:maxLen(params["label"])]
		todo := params["todo"][:maxLen(params["todo"])]
		return apiAdd(label, todo, mainLogger)
	})

	m.Get("/api/count", func() string {
		return apiCount(mainLogger)
	})

	m.Run()
}
