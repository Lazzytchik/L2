package main

import "testing"

func TestGetDataFromFile(t *testing.T) {
	want := []string{
		"man",
		"89",
		"2997",
		"dude hey",
		"how does it",
		"HOW does it",
		"im NOT ready",
		"l2 is fine",
		"36 0 28",
	}
	have, err := getDataFromFile("data.txt")

	if err != nil {
		t.Fatalf("function with error: %s", err)
	}

	if len(have) != len(want) {
		t.Fatalf("test values count mismatch: want %d, got %d.", len(want), len(have))
	}

	for i, s := range have {
		if want[i] != s {
			t.Fatalf("test value mismatch: want %s, got %s", want[i], s)
		}
	}

}
