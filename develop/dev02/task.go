package task02

import (
	"errors"
	"strings"
	"unicode"
)

type Extractor interface {
	Unpack(s string) (string, error)
}

type String struct {
}

func (s *String) Unpack(str string) (string, error) {
	runes := []rune(str)
	answer := strings.Builder{}

	var esc = false
	var prev rune

	for _, v := range runes {
		switch {
		case esc:
			answer.WriteRune(v)
			prev = v
			esc = false
			continue
		case unicode.IsDigit(v):
			if prev == 0 {
				return answer.String(), errors.New("incorrect input string")
			}

			c := int(v) - '0'
			for i := 0; i < c-1; i++ {
				answer.WriteRune(prev)
			}
			continue
		case v == '/':
			esc = true
			continue
		default:
			prev = v
			answer.WriteRune(v)
		}
	}

	return answer.String(), nil
}
