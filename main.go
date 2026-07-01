package main

import (
	"log"

	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmtRegistry := formatter.NewFormatterRegistry()
	fmtRegistry.Register(formatter.NewGoFormatter())

	log.Println("listening for the code format directive on Ctrl-I keypress")
	err := trigger.FormatOnKeyPress(*fmtRegistry)
	if err != nil {
		log.Fatal(err)
	}
}
