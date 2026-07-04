package main

import (
	"context"
	"log"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/config"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"golang.design/x/clipboard"
)

func main() {
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
