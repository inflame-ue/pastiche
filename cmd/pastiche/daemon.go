package main

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
	"golang.design/x/clipboard"
)

func daemon() {
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
}
