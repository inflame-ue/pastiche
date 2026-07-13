package cmd

import (
	"context"
	"log"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/config"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"github.com/kardianos/service"
	"golang.design/x/clipboard"
)

type program struct {
	config   *config.Config
	registry *formatter.FormatterRegistry
	pipeline *pipeline.Pipeline
}

func (p *program) Start(s service.Service) error {
	clipboard.Init()

	ctx := context.Background()
	go p.pipeline.Run(ctx, p.registry)

	log.Println("starting the daemon")
	log.Printf("daemon mode: %s", p.config.Trigger.Mode)

	switch p.config.Trigger.Mode {
	case "autowatch":
		go trigger.FormatAutowatch(ctx, p.pipeline, p.config.Heuristic.Value)
	case "hotkey":
		go trigger.FormatOnKeyPress(p.pipeline, keys.KeyCode(p.config.Hotkey.Key))
	case "both":
		go trigger.FormatAutowatch(ctx, p.pipeline, p.config.Heuristic.Value)
		go trigger.FormatOnKeyPress(p.pipeline, keys.KeyCode(p.config.Hotkey.Key))
	}

	return nil
}

func (p *program) Stop(s service.Service) error {
	log.Println("cleaning up resources and exiting...")
	p.pipeline.Stop()
	return nil
}

func makeService() (service.Service, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, err
	}

	registry := formatter.NewFormatterRegistry()
	registry.Select(conf.Formatters.Order)

	return service.New(&program{
		config:   conf,
		registry: registry,
		pipeline: pipeline.NewPipeline(),
	}, &service.Config{
		Name:        "pastiche",
		DisplayName: "Pastiche",
		Description: "Clipboard code formatter",
		Option: service.KeyValue{
			"UserService": true,
			"RunAtLoad":   true,
			"Restart":     "always", // this is the default, but set explicitly
		},
	})
}
