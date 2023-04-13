package cmd

import (
	"github.com/spf13/cobra"

	"GoTodoCLI/useCase"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new Todo item",
	Long: "Add will create a new todo item to the list",
	Run: useCase.AddRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
