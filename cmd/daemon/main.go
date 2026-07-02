package main

import (
	"context"
	"log"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/formatter/gofmt"
	"github.com/inflame-ue/pastiche/internal/formatter/pythonfmt"
	"github.com/inflame-ue/pastiche/internal/formatter/rustfmt"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
	
	fmtRegistry := formatter.NewFormatterRegistry()
	fmtRegistry.Register(gofmt.NewGoFormatter())
	fmtRegistry.Register(pythonfmt.DefaultPythonFormatter)
	fmtRegistry.Register(rustfmt.DefaultRustFormatter)

	fmtPipeline := pipeline.NewPipeline()
	defer fmtPipeline.Stop()

	log.Println("listening for the code format directive on Ctrl-I keypress")
	go fmtPipeline.Run(context.Background(), fmtRegistry)
	err := trigger.FormatOnKeyPress(fmtPipeline, keys.CtrlI)
	if err != nil {
		log.Fatal(err)
	}
}
