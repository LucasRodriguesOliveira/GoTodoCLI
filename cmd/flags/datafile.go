package flags

import (
	"fmt"
	"os"

	R "GoTodoCLI/cmd/commands"
	v "GoTodoCLI/cmd/variables"
)

func init() {
	R.RootCmd.PersistentFlags().StringVar(
		&v.DatabaseFile.Path,
		"datafile",
		fmt.Sprintf("%s%s", v.HomeDir, string(os.PathSeparator)),
		"data file to store tasks. e.g.: path\\that\\you\\choose",
	)
}
