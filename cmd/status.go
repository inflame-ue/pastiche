package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show configuration and formatter availability",
	Long: `Print the current configuration and check which formatters
are available on the system.

Displays the active trigger mode, hotkey binding, heuristic
threshold, formatter order, and whether each formatter binary
(e.g. black, rustfmt) is installed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
