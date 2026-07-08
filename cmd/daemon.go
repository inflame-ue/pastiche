package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the clipboard formatting daemon",
	Long: `Run the clipboard formatting daemon.

The daemon monitors the system clipboard and formats source code
according to the configured formatters and trigger mode (hotkey,
autowatch, or both).

Requires a config file at ~/.config/pastiche/pastiche.toml.
Run 'pastiche configure' to set up your configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		svc, err := makeService()
		if err != nil {
			log.Fatal(err)
		}

		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}