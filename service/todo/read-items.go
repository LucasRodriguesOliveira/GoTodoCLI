package todo

import (
	to "GoTodoCLI/model/todo"
	fm "GoTodoCLI/service/filemanager"
	"encoding/json"
)

func ReadItems() ([]to.Item, error) {
	fileManager := fm.NewFileManager()
	ok, err := fileManager.Check()

	if err != nil {
		return nil, err
	}

  if !ok {
		return nil, nil
	}

	items, err := fileManager.Read()
	var dbItems []to.Item

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(items, &dbItems); err != nil {
		return nil, err
	}

	for i := range dbItems {
		dbItems[i].SetPosition(i+1)
	}

	return dbItems, nil
}
