package main

import (
	"strings"
)

type Anagramer interface {
	FindAllAnagrams([]string) map[string]string
}

type Russian struct {
}

func (r Russian) FindAllAnagrams(words []string) map[string][]string {
	values := make(map[int64][]string)

	for _, word := range words {
		value := r.countValue(word)

		if v, ok := values[value]; ok {
			values[value] = append(v, word)
		} else {
			values[value] = []string{word}
		}
	}

	answer := make(map[string][]string)

	for _, v := range values {
		if len(v) > 1 {
			answer[v[0]] = v
		}
	}

	return answer
}

func (r Russian) IsAnagram(s1, s2 string) bool {
	return r.countValue(strings.ToLower(s1)) == r.countValue(strings.ToLower(s2))
}

func (r Russian) countValue(s string) int64 {
	runes := []rune(s)
	var count int64

	for _, v := range runes {
		count += 1 << (v - BaseCode)
	}

	return count
}
