package cmd

import (
	"os"

	"github.com/cotramarko/snapvault/internal/engine"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snapvault",
	Short: "A PostgreSQL backup tool for capturing and restoring snapshots of your database.",
	Long: `A PostgreSQL backup tool for capturing and restoring snapshots of your database.

The snapvault CLI tool is intended to be used during local development as an easy way to capture
and restore snapshots of the database, making it possible to quickly restore the database to a
previous state. It supports basic commands such as "save", "restore", "list" and "delete".

The database URL can be specified in multiple ways. Either by a snapvault.toml file
(containing url=<connection-string>), or by setting the environment variable
$DATABASE_URL=<connection-string>, or by passing it as a flag via --url=<connection-string>.
 
The --url flag will always override any of the other ways of specifying the URL. If both a 
snapvault.toml file is present and $DATABASE_URL is set, then the snapvault.toml file will be prioritised.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type Flag string

const (
	URL Flag = "url"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.snapvault.yaml)")
	rootCmd.PersistentFlags().String(string(URL), "", "url to database")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Engine(cmd *cobra.Command) (*engine.Engine, error) {
	url, err := cmd.Flags().GetString(string(URL))
	if err == nil && url != "" {
		return engine.DirectEngine(url)
	}
	path, _ := os.Getwd()
	return engine.LoadEngine(path)
}
