package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Open the interactive configuration TUI",
	Long: `Open the interactive terminal UI to configure pastiche.

Walk through each setting: trigger mode, hotkey binding, heuristic
sensitivity, and formatter order. Writes the config to
~/.config/pastiche/pastiche.toml.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
