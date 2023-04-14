package commands

import (
	"github.com/spf13/cobra"

	"GoTodoCLI/useCase"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do", "d"},
	Short:   "Mark Item as Done",
	Run:     useCase.DoneRun(),
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
