package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	R "GoTodoCLI/cmd/commands"
	V "GoTodoCLI/cmd/variables"
)

// init function is automatically executed, only once, when the package is loaded
// useful when need something to be pre-configured
func init() {
	cobra.OnInitialize(initConfig)
	R.RootCmd.SetVersionTemplate(V.Version)
}

func Execute() {
	if err := R.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
