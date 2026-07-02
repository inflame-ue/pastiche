package pythonfmt

import (
	"fmt"
	"os"
	"os/exec"
)

type PythonFormatter struct {
	formatter string
}

var DefaultPythonFormatter = PythonFormatter{
	formatter: "black",
}

func NewPythonFormatter(formatter string) PythonFormatter {
	return PythonFormatter{
		formatter: formatter,
	}
}

func (pf PythonFormatter) Name() string {
	return "python"
}

func (pf PythonFormatter) Format(src []byte) ([]byte, error) {
	file, err := os.CreateTemp("", "pastiche-*.py")
	if err != nil {
		return nil, fmt.Errorf("creating %s: %w", file.Name(), err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	_, err = file.Write(src)
	if err != nil {
		return nil, fmt.Errorf("writing %s: %w", file.Name(), err)
	}

	cmd := exec.Command(pf.formatter, file.Name())
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("running %s on %s: %w", pf.formatter, file.Name(), err)
	}

	out, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", file.Name(), err)
	}

	return out, nil
}
