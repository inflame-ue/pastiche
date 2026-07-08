package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pastiche",
	Short: "Clipboard code formatter",
	Long: `Pastiche is a clipboard formatting daemon that monitors the
system clipboard and formats source code using pluggable formatters.

Supports Go (gofmt), Python (black), Rust (rustfmt), and more.
Trigger modes: hotkey, autowatch, or both.`,
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
