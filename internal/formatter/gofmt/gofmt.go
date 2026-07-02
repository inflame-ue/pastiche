package gofmt

import (
	"go/format"
)

type goFormatter struct{}

func NewGoFormatter() goFormatter {
	return goFormatter{}
}

func (gf goFormatter) Name() string {
	return "go"
}

func (gf goFormatter) Format(src []byte) ([]byte, error) {
	return format.Source(src)
}
