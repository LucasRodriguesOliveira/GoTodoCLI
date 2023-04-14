package commands

import (
	ut "GoTodoCLI/util"
	v "GoTodoCLI/cmd/variables"

	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "GoTodoCLI",
	Short: "GoTodoCLI is a todo application",
	Version: v.Version,
	Long: `GoTodoCLI will help you to get more done in less time.
	It's designed to be as simple as possible to help you accomplish your goals`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		home := ut.SConcat(
			v.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
		)

		if v.CfgFile == "" {
			v.CfgFile = home
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(
		&v.Db,
		"datafile",
		ut.SConcat(
			v.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
			string(os.PathSeparator),
			"db.json",
		),
		"data file to store tasks. e.g.: path\\that\\you\\choose",
	)

	RootCmd.PersistentFlags().StringVarP(
		&v.CfgFile,
		"config",
		"c",
		ut.SConcat(
			v.HomeDir,
			string(os.PathSeparator),
			"gotodocli",
		),
		"config file",
	)
}
