package main

import "log/syslog"

type TempIndex struct {
	Text string
}

type TempList struct {
	Label string
	Todos []TodoItem
}

func getListValues(label string, mainLogger *syslog.Writer) ([]TodoItem, error) {
	list, err := dbQuery(label)
	if err != nil {
		return nil, err
	}

	tmplList := make([]TodoItem, len(list))
	for i, todo := range list {
		tmplList[i].Id = todo.Id.Hex()
		tmplList[i].Text = todo.Text
	}
	return tmplList, nil
}
