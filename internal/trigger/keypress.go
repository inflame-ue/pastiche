package trigger

import (
	"log"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/inflame-ue/pastiche/internal/formatter"
	"golang.design/x/clipboard"
)

func FormatOnKeyPress(fmtRegistry *formatter.FormatterRegistry, fmtKeyCode keys.KeyCode) error {
	return keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			log.Println("gracefully stopping the keyboard listener...")
			return true, nil
		}

		if key.Code != fmtKeyCode {
			log.Printf("invalid key code, expecting: %d\n", fmtKeyCode)
			return false, nil
		}

		log.Println("formatting the source code in the clipboard...")

		src := clipboard.Read(clipboard.FmtText)
		formatted, err := fmtRegistry.Format(src)
		if err != nil {
			log.Println(err)
			return false, nil
		}

		clipboard.Write(clipboard.FmtText, formatted)

		return false, nil
	})
}
