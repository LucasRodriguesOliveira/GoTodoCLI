package filemanager

import "os"

func (fm *FileManager) create() error {
	file, err := os.Create(fm.File.Name)
	if err != nil {
		return err
	}

	fm.File.Data = file

	return nil
}
