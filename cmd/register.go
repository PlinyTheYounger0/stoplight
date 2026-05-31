/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/PlinyTheYounger0/stoplight/internal/cfg"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "Register a new user",
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Long: `Registers the user in the database and logs the user in as the current user.

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		user, err := programState.Queries.GetUserByName(context.Background(), name)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				newUser, err := programState.Queries.CreateUser(context.Background(), name)
				if err != nil {
					return fmt.Errorf("Error Creating User %s: %v\n", name, err)
				}
				fmt.Printf("User %s Successfully Created\n", newUser.Name)
				if verbose {
					fmt.Printf("User ID: %v\n", user.ID)
					fmt.Printf("Created At: %v\n", user.CreatedAt)
					fmt.Printf("Updated At: %v\n", user.UpdatedAt)
					fmt.Printf("Username: %v\n", user.Name)
				}

				err = cfg.SetCurrentUser(name)
				if err != nil {
					return err
				}

				return nil
			}

			return fmt.Errorf("Error Validating User %s Before Creation: %w\n", name, err)
		}

		return fmt.Errorf("User %s is already registered. Please choose a different name.\n", name)

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
