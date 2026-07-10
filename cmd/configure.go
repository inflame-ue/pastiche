package cmd

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/inflame-ue/pastiche/internal/config"
	"github.com/inflame-ue/pastiche/internal/tui"
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
		cfg, err := config.Load()
		if err != nil {
			cfg = config.NewDefaultConfig()
		}

		p := tea.NewProgram(tui.InitialModel(cfg))
		if _, err := p.Run(); err != nil {
			log.Fatal("the interface has encouneted an unrecoverable error: v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
