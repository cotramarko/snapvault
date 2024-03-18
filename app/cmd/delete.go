/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snapshots := internal.GetSnapshots()
		name := args[0]
		if snapshots.HasSnapshotWithName(name) {
			fmt.Printf("Deleting snapshot `%s`\n", name)
		} else {
			fmt.Printf("Snapshot `%s` does not exist\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
