package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	V "GoTodoCLI/cmd/variables"
)

func initConfig() {
	viper.AddConfigPath(V.CfgFile)
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}
