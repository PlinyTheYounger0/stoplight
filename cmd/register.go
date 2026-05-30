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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Long: `Registers the user in the database and logs the user in as the current user.

`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		user, err := programState.DB.CreateUser(context.Background(), name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Creating User: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("User %s Successfully Created\n", user.Name)

		if verbose {
			fmt.Printf("User ID: %v\n", user.ID)
			fmt.Printf("Created At: %v\n", user.CreatedAt)
			fmt.Printf("Updated At: %v\n", user.UpdatedAt)
			fmt.Printf("Username: %v\n", user.Name)
		}
	},
}

func init() {
	registerCmd.Flags().BoolP("verbose", "v", false, "enable verbose output")
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
