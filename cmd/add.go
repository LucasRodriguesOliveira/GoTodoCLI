package cmd

import (
	"github.com/spf13/cobra"

	"GoTodoCLI/useCase"
)

var priority int

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new Todo item",
	Long:    "Add will create a new todo item to the list",
	Run:     useCase.AddRun(
		&DatabaseFile,
		&priority,
	),
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
