package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/config"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
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
		err := clipboard.Init()
		if err != nil {
			log.Fatal(err)
		}

		conf, err := config.Load()
		if err != nil {
			log.Fatal(err)
		}

		fmtRegistry := formatter.NewFormatterRegistry()
		fmtRegistry.Select(conf.Formatters.Order)

		fmtPipeline := pipeline.NewPipeline()
		ctx := context.Background()
		go fmtPipeline.Run(ctx, fmtRegistry)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			log.Print("cleaning up resources and exiting...")
			fmtPipeline.Stop()
			os.Exit(0)
		}()

		log.Printf("starting the daemon...")
		log.Printf("daemon mode: %s", conf.Trigger.Mode)
		switch conf.Trigger.Mode {
		case "autowatch":
			trigger.FormatAutowatch(ctx, fmtPipeline, conf.Heuristic.Value)
		case "hotkey":
			trigger.FormatOnKeyPress(fmtPipeline, keys.KeyCode(conf.Hotkey.Key))
		case "both":
			go trigger.FormatAutowatch(ctx, fmtPipeline, conf.Heuristic.Value)
			trigger.FormatOnKeyPress(fmtPipeline, keys.KeyCode(conf.Hotkey.Key))
		default:
			log.Fatal("unknown daemon mode")
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
