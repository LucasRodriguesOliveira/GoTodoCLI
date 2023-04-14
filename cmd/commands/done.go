package commands

import (
	"github.com/spf13/cobra"

	"GoTodoCLI/useCase"
	V "GoTodoCLI/cmd/variables"
)

var doneCmd = &cobra.Command{
	Use: "done",
	Aliases: []string{"do", "d"},
	Short: "Mark Item as Done",
	Run: useCase.DoneRun(&V.DatabaseFile),
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
