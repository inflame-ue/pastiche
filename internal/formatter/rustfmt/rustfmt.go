package rustfmt

import (
	"fmt"
	"os"
	"os/exec"
)

type rustFormatter struct {
	formatter string
}

var DefaultRustFormatter = rustFormatter{
	formatter: "rustfmt",
}

func NewRustFormatter(formatter string) rustFormatter {
	return rustFormatter{
		formatter: formatter,
	}
}

func (rf rustFormatter) Name() string {
	return "rust"
}

func (rf rustFormatter) Format(src []byte) ([]byte, error) {
	file, err := os.CreateTemp("", "pastiche-*.rs")
	if err != nil {
		return nil, fmt.Errorf("creating %s: %w", file.Name(), err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	_, err = file.Write(src)
	if err != nil {
		return nil, fmt.Errorf("writing %s: %w", file.Name(), err)
	}

	cmd := exec.Command(rf.formatter, file.Name(), "--force")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("running %s on %s: %w", rf.formatter, file.Name(), err)
	}

	out, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", file.Name(), err)
	}

	return out, nil
}
