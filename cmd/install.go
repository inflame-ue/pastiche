package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install pastiche as a systemd user service",
	Long: `Install pastiche as a systemd user service so it starts
automatically on login.

Creates a systemd unit file at ~/.config/systemd/user/pastiche.service
pointing at the current binary and enables it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
