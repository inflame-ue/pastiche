package formatter

import (
	"errors"
	"log"
)

type Formatter interface {
	Name() string
	Format(src []byte) ([]byte, error)
}

type FormatterRegistry struct {
	formatters []Formatter
}

func NewFormatterRegistry() *FormatterRegistry {
	return &FormatterRegistry{
		formatters: []Formatter{},
	}
}

func (fr *FormatterRegistry) Register(formatter Formatter) {
	fr.formatters = append(fr.formatters, formatter)
}

func (fr *FormatterRegistry) Format(src []byte) ([]byte, error) {
	for _, formatter := range fr.formatters {
		out, err := formatter.Format(src)
		if err == nil {
			return out, nil
		}
		log.Print(err)
	}
	return nil, errors.New("no formatter could parse the source")
}
