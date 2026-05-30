/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// usersResetCmd represents the usersReset command
var usersResetCmd = &cobra.Command{
	Use:   "usersReset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := programState.Queries.ResetUsers(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Resetting Users DB: %v\n", err)
		}

		fmt.Println("Users DB Successfully Reset")
	},
}

func init() {
	rootCmd.AddCommand(usersResetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersResetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersResetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
