package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Read In Config: %v", err.Error())
		os.Exit(1)
	}
}
