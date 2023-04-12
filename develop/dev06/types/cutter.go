package cut

import (
	"strings"
)

type Cutter interface {
	Cut(s string) (string, bool)
}

type Field struct {
	Index     int
	Delimiter string
}

func (f Field) Cut(s string) (string, bool) {
	split := strings.Split(s, f.Delimiter)
	sep := false

	if len(split) > 0 {
		sep = true
	}

	if len(split) < f.Index {
		return s, sep
	}

	return strings.Split(s, f.Delimiter)[f.Index-1], true
}

type Default struct {
}

func (d Default) Cut(s string) (string, bool) {
	return s, true
}
