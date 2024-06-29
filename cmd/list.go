package cmd

import (
	"fmt"
	"os"

	"github.com/cotramarko/snapvault/internal/commands"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List snapshots",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		e, err := Engine(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		snapshots, err := commands.List(*e)

		if err != nil {
			fmt.Println(err)
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Created", "Size"})
		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: "Created", AlignHeader: text.AlignCenter, AlignFooter: text.AlignRight},
			{Name: "Size", AlignHeader: text.AlignRight, Align: text.AlignRight},
		})

		for _, d := range snapshots {
			t.AppendRow(table.Row{d.SnapName, d.Created, d.Size})
		}
		t.AppendSeparator()
		/*
			t.AppendFooter(table.Row{
				"",
				"Total",
				fmt.Sprintf("%d MB", 0),
			})
		*/
		t.SetStyle(table.StyleRounded)
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
