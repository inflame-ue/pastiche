package trigger

import (
	"log"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"golang.design/x/clipboard"
)

func FormatOnKeyPress(fmtRegistry formatter.FormatterRegistry) error {
	return keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			log.Println("stopping the keyboard listener...")
			return true, nil
		}

		if key.Code != keys.CtrlI {
			return false, nil
		}

		log.Println("formatting the code in the clipboard...")

		src := clipboard.Read(clipboard.FmtText)
		formatted, err := fmtRegistry.Format(src)
		if err != nil {
			return true, err
		}

		clipboard.Write(clipboard.FmtText, formatted)

		return false, nil
	})
}
