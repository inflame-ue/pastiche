package formatter

import (
	"go/format"
)

type GoFormatter struct{}

func NewGoFormatter() GoFormatter {
	return GoFormatter{}
}

func (gf GoFormatter) Name() string {
	return "go"
}

func (gf GoFormatter) Format(src []byte) ([]byte, error) {
	return format.Source(src)
}
