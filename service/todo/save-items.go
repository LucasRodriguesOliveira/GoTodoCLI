package todo

import (
	model "GoTodoCLI/model/todo"
	"encoding/json"

	fm "GoTodoCLI/service/filemanager"
)

func SaveItems(items []model.Item) error {
	fileManager := fm.NewFileManager()
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
