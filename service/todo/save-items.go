package todo

import (
	"GoTodoCLI/model"
	"encoding/json"

	fm "GoTodoCLI/service/filemanager"
)

func SaveItems(items []model.Item) error {
	fileManager := fm.NewFileManager("db.json", ".\\")
	
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
