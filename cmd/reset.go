/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var tables []string
// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset database tables",
	Long: `Reset is used to reset the database. This is a command used for testing purposes and not meant to be utilized in production.
	You can specify which table specificly you want dropped or you can provide no flags and it will drop all tables.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(tables) == 0 {
			err := programState.Queries.ResetDB(context.Background())
			if err != nil {
				return fmt.Errorf("Error Resetting Database: %w\n", err)
			}
		}

		resets := map[string]func(context.Context) error {
			"users": programState.Queries.ResetUsers,
		}

		for _, table := range tables {
			fn, ok := resets[table]
			if !ok {
				return fmt.Errorf("Unknown Table: %s\n", table)
			}

			if err := fn(context.Background()); err != nil {
				return fmt.Errorf("Error Resetting %s: %w\n", table, err)
			}

		}

		return nil
	},
}

func init() {
	resetCmd.Flags().StringSliceVarP(&tables, "table", "t", nil, "Select a specific table to reset")
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
