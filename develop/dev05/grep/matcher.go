package grep

import "strings"

type Matcher interface {
	Match(s, match string) bool
}

type Have struct {
}

func (h Have) Match(s, match string) bool {
	return strings.Contains(s, match)
}

type Exact struct {
}

func (e Exact) Match(s, match string) bool {
	return s == match
}
