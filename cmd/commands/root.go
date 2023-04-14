package commands

import (
	C "GoTodoCLI/cmd/constants"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "GoTodoCLI",
	Short: "GoTodoCLI is a todo application",
	Version: C.Version,
	Long: `GoTodoCLI will help you to get more done in less time.
	It's designed to be as simple as possible to help you accomplish your goals`,
}
