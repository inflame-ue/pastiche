package formatter

import (
	"errors"
	"log"

	"github.com/inflame-ue/pastiche/internal/formatter/gofmt"
	"github.com/inflame-ue/pastiche/internal/formatter/pythonfmt"
	"github.com/inflame-ue/pastiche/internal/formatter/rustfmt"
)

type Formatter interface {
	Name() string
	Format(src []byte) ([]byte, error)
}

type FormatterRegistry struct {
	formatters []Formatter
	lookup     map[string]Formatter
}

func NewFormatterRegistry() *FormatterRegistry {
	reg := &FormatterRegistry{
		formatters: []Formatter{},
		lookup:     make(map[string]Formatter),
	}
	reg.Register(gofmt.NewGoFormatter())
	reg.Register(pythonfmt.DefaultPythonFormatter)
	reg.Register(rustfmt.DefaultRustFormatter)
	return reg
}

func (fr *FormatterRegistry) Register(formatter Formatter) {
	fr.lookup[formatter.Name()] = formatter
}

func (fr *FormatterRegistry) Select(names []string) {
	for _, name := range names {
		formatter, ok := fr.lookup[name]
		if !ok {
			log.Print("no formatter with this name available")
			continue
		}
		fr.formatters = append(fr.formatters, formatter)
	}
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
