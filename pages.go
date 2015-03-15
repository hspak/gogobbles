package main

import "gopkg.in/mgo.v2"

type TempIndex struct {
	ListCount string
	AllLists  map[string]int
}

type TempList struct {
	Label string
	Todos []TodoItem
}

func getIndexInfo(session *mgo.Session) (map[string]int, int, error) {
	lists, count, err := dbCountLists(session)
	if err != nil {
		return nil, 0, err
	}
	return lists, count, nil
}

func getListValues(session *mgo.Session, label string) ([]TodoItem, error) {
	list, err := dbQuery(session, label)
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
