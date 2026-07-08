package cmd

import (
	"fmt"
	"os/exec"

	"github.com/inflame-ue/pastiche/internal/config"
	"github.com/inflame-ue/pastiche/internal/tui"
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
		conf, err := config.Load()
		if err != nil {
			fmt.Println("no config found at ~/.config/pastiche/pastiche.toml")
			fmt.Println("run 'pastiche configure' to set up")
			return
		}

		fmt.Println("Trigger mode:", conf.Trigger.Mode)
		if k := tui.KeyName(conf.Hotkey.Key); k != "" {
			fmt.Println("Hotkey:      ", k)
		}
		fmt.Println("Heuristic threshold:", conf.Heuristic.Value)
		fmt.Println()
		fmt.Println("Formatters (in order):")
		for _, f := range conf.Formatters.Order {
			path, err := exec.LookPath(f)
			if err == nil {
				fmt.Printf("  ✓ %s (%s)\n", f, path)
			} else {
				fmt.Printf("  ✗ %s (not found)\n", f)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}