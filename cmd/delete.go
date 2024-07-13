package cmd

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal/commands"
	"github.com/cotramarko/snapvault/internal/engine"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		e, err := Engine(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		snapName := engine.SnapName(args[0])
		err = commands.Delete(*e, snapName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Deleted snapshot", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
