package main

import (
	"context"
	"log"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"golang.design/x/clipboard"
)

func init() {
	// this is fine, because we don't steal error agency from the caller
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmtRegistry := formatter.NewFormatterRegistry()
	fmtRegistry.Register(formatter.NewGoFormatter())
	fmtPipeline := pipeline.NewPipeline()
	defer fmtPipeline.Stop()

	log.Println("listening for the code format directive on Ctrl-I keypress")
	go fmtPipeline.Run(context.Background(), fmtRegistry)
	err := trigger.FormatOnKeyPress(fmtPipeline, keys.CtrlI)
	if err != nil {
		log.Fatal(err)
	}
}
