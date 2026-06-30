package formatter

import "errors"

type Formatter interface {
	Name() string
	Format(src []byte) ([]byte, error)
}

type FormatterRegistry struct {
	formatters []Formatter
}

func (fr *FormatterRegistry) Format(src []byte) ([]byte, error) {
	for _, formatter := range fr.formatters {
		out, err := formatter.Format(src)
		if err == nil {
			return out, nil
		}
	}
	return nil, errors.New("err: no formatter could parse the source")
}
