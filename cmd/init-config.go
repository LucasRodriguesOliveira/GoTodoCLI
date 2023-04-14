package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	V "GoTodoCLI/cmd/variables"
	R "GoTodoCLI/cmd/commands"
)

func initConfig() {
	viper.AddConfigPath(V.CfgFile)
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %q\n\n", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}

	viper.BindPFlag("db", R.RootCmd.PersistentFlags().Lookup("datafile"))
}
