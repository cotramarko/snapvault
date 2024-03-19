/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saveCmd represents the backup command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Create a snapshot",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating snapshot", args[0])
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
