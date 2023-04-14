package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"GoTodoCLI/useCase"
	V "GoTodoCLI/cmd/variables"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: fmt.Sprintf(
		"lists all the items stored in \"%s%s%s\"",
		V.DatabaseFile.Path,
		string(os.PathSeparator),
		V.DatabaseFile.Name,
	),
	Long: "List will look for the storage file and list all the items stored",
	Run: useCase.ListRun(&V.DatabaseFile),
}

func init() {
	RootCmd.AddCommand(listCmd)
}
