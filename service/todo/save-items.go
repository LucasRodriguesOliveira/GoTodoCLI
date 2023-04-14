package todo

import (
	model "GoTodoCLI/model/todo"
	"encoding/json"

	fm "GoTodoCLI/service/filemanager"
	t  "GoTodoCLI/service/types"
)

func SaveItems(items []model.Item, dbFile t.FileData) error {
	fileManager := fm.NewFileManager(dbFile)
	ok, err := fileManager.Check()

	if err != nil {
		return err
	}

	if ok {
		data, err := fileManager.Read()

		if err != nil {
			return err
		}

		var dbItems []model.Item

		if err = json.Unmarshal(data, &dbItems); err != nil {
			return err
		}

		items = append(dbItems, items...)
	}

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
