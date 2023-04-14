package cmd

import (
	c "GoTodoCLI/cmd/config"
	t "GoTodoCLI/service/types"
)

var (
	homeDir string = c.HomeDir()
	DatabaseFile = t.DatabaseFile{Name: dbFileName}
)
