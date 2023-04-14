package config

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
)

func HomeDir() string {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println("Could not find the home directory. Please set data file directory using --datafile")
		panic(1)
	}

	return home
}
