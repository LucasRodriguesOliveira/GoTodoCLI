package commands

import (
	C "GoTodoCLI/cmd/constants"
	V "GoTodoCLI/cmd/variables"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "GoTodoCLI",
	Short: "GoTodoCLI is a todo application",
	Version: C.Version,
	Long: `GoTodoCLI will help you to get more done in less time.
	It's designed to be as simple as possible to help you accomplish your goals`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if V.DatabaseFile.Path == "" {
			V.DatabaseFile.Path = fmt.Sprintf("%s%s", V.HomeDir, string(os.PathSeparator))
		}
	},
}
