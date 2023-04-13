package filemanager

import (
	"fmt"
	"os"
)

func (fm *FileManager) Read() ([]byte, error) {
  ok, err := fm.Check()

	if err != nil {
		return nil, err
	}

  if !ok {
		return nil, fmt.Errorf("file %q does not exists", fm.File.Name)
	}

	data, err := os.ReadFile(fm.File.Name)

	if err != nil {
		return nil, err
	}

	return data, nil
}
