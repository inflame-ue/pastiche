package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install pastiche as a background service",
	Long: `Install pastiche as a background service so it starts
automatically on login (Linux: systemd user service,
macOS: launchd, Windows: Windows Service).

Run 'pastiche daemon' directly to start without installing.`,
	Run: func(cmd *cobra.Command, args []string) {
		svc, err := makeService()
		if err != nil {
			log.Fatal(err)
		}

		if err := svc.Install(); err != nil {
			log.Printf("service already installed, reinstalling...")
			svc.Stop()
			svc.Uninstall()
			if err := svc.Install(); err != nil {
				log.Fatalf("install failed: %v", err)
			}
		}
		log.Println("service installed")

		if err := svc.Start(); err != nil {
			log.Fatalf("start failed: %v", err)
		}
		log.Println("service started")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}