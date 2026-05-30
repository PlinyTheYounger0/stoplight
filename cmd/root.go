/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/PlinyTheYounger0/stoplight/internal/cfg"
	"github.com/PlinyTheYounger0/stoplight/internal/database"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

var programState *cfg.State

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stoplight",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	 PersistentPreRunE: func(cmd *cobra.Command, args []string) error { 

		if programState == nil {

			db, err := sql.Open("postgres", viper.GetString("db_url"))
			if err != nil {
				return fmt.Errorf("Error Establishing DB Connection: %v\n", err.Error())
			}

			if verbose {
				fmt.Printf("Successfully Connected to DB: %s\n", viper.GetString("db_url"))
			}

			dbQueries := database.New(db)

			programState = &cfg.State{
				Queries: dbQueries,
				DB: db,
			}
		}

		return nil
	 },

	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if programState != nil && programState.DB != nil {
			if err := programState.DB.Close(); err != nil {
				return fmt.Errorf("Error Closing DB Connection: %v\n", err.Error())
			}
		}

		return nil
	}, 

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/stoplight/internal/config/.stoplight.env)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("json")
		viper.SetConfigName(".stoplight")
		viper.AddConfigPath("$HOME/stoplight/internal/cfg/")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
