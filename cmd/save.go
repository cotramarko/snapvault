/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal/commands"
	"github.com/cotramarko/snapvault/internal/config"
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
		engine := config.GetDefaultEngine()
		err := commands.Save(*engine, args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
