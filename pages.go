package main

type TempIndex struct {
	ListCount string
}

type TempList struct {
	Label string
	Todos []TodoItem
}

func getIndexInfo() (int, error) {
	_, count, err := dbCountLists()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func getListValues(label string) ([]TodoItem, error) {
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
