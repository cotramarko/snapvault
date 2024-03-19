/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cotramarko/snapvault/internal/commands"
	"github.com/cotramarko/snapvault/internal/config"
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
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Created", "Size"})
		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: "Created", AlignHeader: text.AlignCenter, AlignFooter: text.AlignRight},
			{Name: "Size", AlignHeader: text.AlignRight, Align: text.AlignRight},
		})

		engine := config.GetDefaultEngine()
		snapshots, err := commands.List(*engine)

		if err != nil {
			panic(err)
		}
		for _, d := range snapshots {
			t.AppendRow(table.Row{d, time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("%d MB", 0)})
		}
		t.AppendSeparator()
		t.AppendFooter(table.Row{
			"",
			"Total",
			fmt.Sprintf("%d MB", 0),
		})
		t.SetStyle(table.StyleRounded)
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
