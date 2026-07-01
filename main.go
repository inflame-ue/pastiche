package main

import (
	"log"

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
	log.Println("listening for the code format directive on Ctrl-I keypress")
	err := trigger.FormatOnKeyPress()
	if err != nil {
		log.Fatal(err)
	}
}
