package cmd

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVar(
		&DatabaseFile.Path,
		"datafile",
		fmt.Sprintf("%s%s", homeDir, string(os.PathSeparator)),
		"data file to store tasks. e.g.: path\\that\\you\\choose",
	)
}
