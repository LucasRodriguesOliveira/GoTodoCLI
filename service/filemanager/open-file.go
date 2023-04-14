package filemanager

import (
	"os"
)

func (fm *FileManager) open() error {
  file, err := os.OpenFile(fm.File.Name, os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	fm.File.Data = file

	return nil
}
