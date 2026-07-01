package trigger

import (
	"log"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/pipeline"
	"golang.design/x/clipboard"
)

func FormatOnKeyPress(p *pipeline.Pipeline, keyCode keys.KeyCode) error {
	return keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			log.Println("stopping the keyboard listener...")
			return true, nil
		}

		if key.Code != keyCode {
			log.Printf("invalid key code, expecting: %d\n", keyCode)
			return false, nil
		}

		log.Println("sending the source code in the clipboard to the formatting pipeline...")
		p.Submit(clipboard.Read(clipboard.FmtText))
		return false, nil
	})
}
