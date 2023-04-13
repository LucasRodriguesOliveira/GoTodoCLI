package filemanager

import (
	"fmt"
	"os"
)

type fileData struct {
	Name string
	Data *os.File
}

type FileManager struct {
  File fileData
}

func NewFileManager(fileName, path string) *FileManager {
	return &FileManager{
		File: fileData{
			Name: fmt.Sprintf("%s%s", path, fileName),
		},
	}
}
