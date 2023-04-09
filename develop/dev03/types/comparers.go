package dev03

import (
	"unicode"
)

type Comparer[T comparable] interface {
	Less(T, T) bool
}

type Alphabet struct {
}

func (sorter Alphabet) Less(i string, j string) bool {
	irunes := []rune(i)
	jrunes := []rune(j)

	max := len(i)
	if max > len(j) {
		max = len(j)
	}

	for ix := 0; ix < max; ix++ {
		ir := irunes[ix]
		jr := jrunes[ix]

		if unicode.ToLower(ir) != unicode.ToLower(jr) {
			return ir < jr
		}

		if unicode.IsUpper(ir) != unicode.IsUpper(jr) {
			return unicode.IsLower(ir)
		}
	}

	return len(i) < len(j)
}
