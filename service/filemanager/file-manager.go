package filemanager

import (
	"fmt"
	"os"

	t "GoTodoCLI/service/types"
)

type fileData struct {
	Name string
	Data *os.File
}

type FileManager struct {
	File fileData
}

func NewFileManager(f t.FileData) *FileManager {
	fileName, path := f.Get()

	return &FileManager{
		File: fileData{
			Name: fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), fileName),
		},
	}
}
