package pythonfmt

import (
	"fmt"
	"os"
)

type PythonFormatter struct{}

func NewPythonFormatter() PythonFormatter {
	return PythonFormatter{}
}

func (pf PythonFormatter) Name() string {
	return "python"
}

func (pf PythonFormatter) Format(src []byte) ([]byte, error) {
	err := os.WriteFile("temporary.py", src, os.ModeTemporary)
	if err != nil {
		return nil, fmt.Errorf("writing temporary python file: %w", err)
	}

	return []byte{}, nil
}
