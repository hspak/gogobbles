package main

import (
	"fmt"
	"log"
	"log/syslog"
	"strconv"

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
		count, err := getIndexInfo()
		fmt.Println(count)
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			r.HTML(500, "index", nil) // make a new tmpl for this
		}
		r.HTML(200, "index", TempIndex{ListCount: strconv.Itoa(count)})
	})

	m.Get("/list/:label", func(params martini.Params, r render.Render) {
		tmplList, err := getListValues(params["label"])
		if err != nil {
			mainLogger.Err("Error: db query went bad: " + err.Error())
			r.HTML(500, "index", nil) // make a new tmpl for this
		}
		r.HTML(200, "list", TempList{Label: params["label"], Todos: tmplList})
	})

	// API
	m.Get("/api/get/:label", func(params martini.Params) string {
		return apiGet(params["label"][:80], params["todo"][:80], mainLogger)
	})

	m.Get("/api/remove/:label/:id", func(params martini.Params) string {
		return apiRemove(params["label"][:80], params["id"], mainLogger)
	})

	m.Get("/api/add/:label/:todo", func(params martini.Params) string {
		return apiAdd(params["label"][:80], params["todo"][:80], mainLogger)
	})

	m.Get("/api/count", func() string {
		return apiCount(mainLogger)
	})

	m.Run()
}
