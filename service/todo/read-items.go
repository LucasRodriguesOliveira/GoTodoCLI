package todo

import (
	to "GoTodoCLI/model/todo"
	fm "GoTodoCLI/service/filemanager"
	t  "GoTodoCLI/service/types"
	"encoding/json"
	"fmt"
)

func ReadItems(dbFile t.FileData) ([]to.Item, error) {
	fileManager := fm.NewFileManager(dbFile)
	ok, err := fileManager.Check()

	if err != nil {
		return nil, err
	}

  if !ok {
		fmt.Println("No task's found, use command <add> to add a task or more.")
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
