package cmd

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal/commands"
	"github.com/cotramarko/snapvault/internal/engine"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a snapshot",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		e, err := Engine(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		snapName := engine.SnapName(args[0])
		err = commands.Restore(*e, snapName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Restored snapshot", args[0])
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

}
