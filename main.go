package main

import (
	"log"

	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"github.com/inflame-ue/pastiche/internal/trigger"
	"golang.design/x/clipboard"
)

func init() {
	// this is fine, because we don't steal error agency from the client
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmtRegistry := formatter.NewFormatterRegistry()
	fmtRegistry.Register(formatter.NewGoFormatter())

	log.Println("listening for the code format directive on Ctrl-I keypress")
	err := trigger.FormatOnKeyPress(fmtRegistry, keys.CtrlI)
	if err != nil {
		log.Fatal(err)
	}
}
