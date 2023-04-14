package filemanager

import (
	"os"

	"github.com/spf13/viper"
)

type fileData struct {
	Name string
	Data *os.File
}

type FileManager struct {
	File fileData
}

func NewFileManager() *FileManager {
	return &FileManager{
		File: fileData{
			Name: viper.GetString("db"),
		},
	}
}
