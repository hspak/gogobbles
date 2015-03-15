package main

import "gopkg.in/mgo.v2"

type TempIndex struct {
	ListCount string
}

type TempList struct {
	Label string
	Todos []TodoItem
}

func getIndexInfo(session *mgo.Session) (int, error) {
	_, count, err := dbCountLists(session)
	if err != nil {
		return 0, err
	}

	return count, nil
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
