package dev03

import (
	"strconv"
	"strings"
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

type Numeric struct {
}

func (n *Numeric) Less(i string, j string) bool {
	iNumber := n.extractNumber(i)
	jNumber := n.extractNumber(j)

	if len(i) != len(j) {
		return len(iNumber) < len(jNumber)
	}

	for ix := 0; ix < len(iNumber); ix++ {
		ir := rune(iNumber[ix])
		jr := rune(jNumber[ix])

		if ir != jr {
			return ir < jr
		}
	}

	return len(i) < len(j)
}

func (n *Numeric) extractNumber(s string) string {
	answer := strings.Builder{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			answer.WriteRune(c)
		} else {
			return answer.String()
		}
	}

	return answer.String()
}

type Month struct {
	months   map[string]int
	Alphabet Alphabet
}

func (m *Month) Init() *Month {
	m.SetEngMonths()
	m.Alphabet = Alphabet{}
	return m
}

func (m *Month) SetEngMonths() {
	m.months = map[string]int{
		"JAN": 1,
		"FEB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}
}

func (m *Month) ExtractMonth(s string) int {
	if v, ok := m.months[strings.ToUpper(s)[:3]]; ok {
		return v
	}
	return 0
}

func (m *Month) Less(i string, j string) bool {
	iMonth := m.ExtractMonth(i)
	jMonth := m.ExtractMonth(j)

	if iMonth != 0 && jMonth != 0 {
		return iMonth < jMonth
	} else if iMonth != 0 || jMonth != 0 {
		return iMonth != 0
	}

	return m.Alphabet.Less(i, j)
}

type Bytes struct {
	suf map[rune]int
}

func (b *Bytes) SetBaseSuffixes() {
	b.suf = map[rune]int{
		'K': 1024,
		'M': 1024 * 1024,
		'G': 1024 * 1024 * 1024,
	}
}

func (b *Bytes) Less(i string, j string) bool {
	iNumber := b.extractNumber(i)
	jNumber := b.extractNumber(j)

	return iNumber < jNumber
}

func (b *Bytes) extractNumber(s string) int {
	answer := strings.Builder{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			answer.WriteRune(c)
		} else if v, ok := b.suf[c]; ok {
			number, _ := strconv.Atoi(answer.String())
			return number * v
		} else {
			number, _ := strconv.Atoi(answer.String())
			return number
		}
	}

	number, _ := strconv.Atoi(answer.String())
	return number
}
