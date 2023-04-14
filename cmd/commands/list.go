package commands

import (
	"fmt"

	T "GoTodoCLI/cmd/types"
	"GoTodoCLI/useCase"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showDone bool = false
	showAll  bool = false

	f T.FBool
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: fmt.Sprintf(
		"lists all the items stored in %q",
		viper.GetString("db"),
	),
	Long: "List will look for the storage file and list all the items stored",
	Run: useCase.ListRun(&f),
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
