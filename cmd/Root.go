package cmd

import (
	"fmt"
	"os"

	// homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

const (
	version string = "v0.0.1"
)

var rootCmd = &cobra.Command{
	Use:   "GoTodoCLI",
	Short: "GoTodoCLI is a todo application",
	Version: version,
	Long: `GoTodoCLI will help you to get more done in less time.
	It's designed to be as simple as possible to help you accomplish your goals`,
}

// var cfgFile string = ""

// init function is automatically executed, only once, when the package is loaded
// useful when need something to be pre-configured
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(version)
}

func initConfig() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
