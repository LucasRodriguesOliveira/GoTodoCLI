package todo

import (
	model "GoTodoCLI/model/todo"
	"encoding/json"

	fm "GoTodoCLI/service/filemanager"
	t  "GoTodoCLI/service/types"
)

func SaveItems(items []model.Item, dbFile t.FileData) error {
	fileManager := fm.NewFileManager(dbFile)
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = fileManager.Write(b)

	if err != nil {
		return err
	}

	return nil
}
