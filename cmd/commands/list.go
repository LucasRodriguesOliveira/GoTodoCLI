package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"GoTodoCLI/useCase"
	V "GoTodoCLI/cmd/variables"
	T "GoTodoCLI/cmd/types"
)

var (
	showDone bool = false
	showAll  bool = false

	f T.FBool
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
	Run: useCase.ListRun(&V.DatabaseFile, &f),
	PreRun: func(cmd *cobra.Command, args []string) {
		f = make(map[string]bool)
		f["done"] = showDone
		f["all"] = showAll
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&showDone, "done", "d", false, "Show 'Done' tasks")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
}
