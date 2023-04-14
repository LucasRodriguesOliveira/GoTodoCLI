package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"GoTodoCLI/useCase"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: fmt.Sprintf(
		"lists all the items stored in \"%s%s%s\"",
		DatabaseFile.Path,
		string(os.PathSeparator),
		DatabaseFile.Name,
	),
	Long: "List will look for the storage file and list all the items stored",
	Run: useCase.ListRun(&DatabaseFile),
}

func init() {
	rootCmd.AddCommand(listCmd)
}
