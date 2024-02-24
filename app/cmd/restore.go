/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal/restore"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a snapshot",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snapshots := restore.GetSnapshots()
		if snapshots.HasSnapshotWithName(args[0]) {
			fmt.Println("Restoring snapshot", args[0])
		} else {
			fmt.Println("Snapshot", args[0], "does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

}
