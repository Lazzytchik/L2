package main

import (
	"testing"
)

func TestRussian_IsAnagram(t *testing.T) {
	want := false

	s1 := "авв"
	s2 := "абд"

	a := Russian{}

	if a.IsAnagram(s1, s2) != want {
		t.Fatalf("wrong answer: %s must%s be an anagram for %s", s1, not(!want), s2)
	}
}

func TestRussian_IsAnagram_Negative(t *testing.T) {
	want := false

	s1 := "тпбур"
	s2 := "траур"

	a := Russian{}

	if a.IsAnagram(s1, s2) != want {
		t.Fatalf("wrong answer: %s must%s be an anagram for %s", s1, not(!want), s2)
	}
}

func TestRussian_IsAnagram_Positive(t *testing.T) {
	want := true

	s1 := "тяпка"
	s2 := "пятка"

	a := Russian{}

	if a.IsAnagram(s1, s2) != want {
		t.Fatalf("wrong answer: %s must%s be an anagram for %s", s1, not(!want), s2)
	}
}

func TestRussian_IsAnagram_PositiveUpperChars(t *testing.T) {
	want := true

	s1 := "тЯпка"
	s2 := "пятКа"

	a := Russian{}

	if a.IsAnagram(s1, s2) != want {
		t.Fatalf("wrong answer: %s must%s be an anagram for %s", s1, not(!want), s2)
	}
}

func TestRussian_FindAllAnagrams(t *testing.T) {
	input := []string{
		"тяпка",
		"слиток",
		"тряпка",
		"столик",
		"листок",
		"пятка",
	}

	want := map[string][]string{
		"тяпка":  {"тяпка", "пятка"},
		"слиток": {"слиток", "столик", "листок"},
	}

	a := Russian{}

	result := a.FindAllAnagrams(input)

	if !AreMapsEqual(result, want) {
		t.Fatalf("wrong answer: want %s, got %s", want, result)
	}
}

func not(not bool) string {
	if not {
		return " not"
	}

	return ""
}

func AreMapsEqual(m1, m2 map[string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, strings := range m1 {
		v, ok := m2[key]
		if !ok {
			return false
		}

		if !AreSlicesEqual(strings, v) {
			return false
		}
	}

	return true
}

func AreSlicesEqual(s1, s2 []string) bool {
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
