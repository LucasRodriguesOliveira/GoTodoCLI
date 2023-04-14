package commands

import (
	C "GoTodoCLI/cmd/constants"
	V "GoTodoCLI/cmd/variables"
	ut "GoTodoCLI/util"

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
		home := ut.SConcat(
			V.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
		)

		if V.DatabaseFile.Path == "" {
			V.DatabaseFile.Path = home
		}

		if V.CfgFile == "" {
			V.CfgFile = home
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(
		&V.DatabaseFile.Path,
		"datafile",
		ut.SConcat(
			V.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
		),
		"data file to store tasks. e.g.: path\\that\\you\\choose",
	)

	RootCmd.PersistentFlags().StringVarP(
		&V.CfgFile,
		"config",
		"c",
		ut.SConcat(
			V.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
		),
		"config file",
	)
}
