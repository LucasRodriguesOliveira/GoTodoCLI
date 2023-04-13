package filemanager

import (
	"fmt"
	"os"
)

/*
*
	Returns true, if the specified file in *FileManager exists
*/
func (fm *FileManager) Check() (bool, error) {
	info, err := os.Stat(fm.File.Name)

	if err == nil && info.IsDir() {
		return false, fmt.Errorf("%q is a directory", info.Name())
	}

	return !os.IsNotExist(err), nil
}
