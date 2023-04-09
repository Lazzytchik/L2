package dev03

import "testing"

func TestAlphabet_Less_PositiveUpper(t *testing.T) {
	lesser := "abc"
	bigger := "Abc"

	comp := Alphabet{}

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}

func TestAlphabet_Less_PositiveNumber(t *testing.T) {
	lesser := "8abc"
	bigger := "abc"

	comp := Alphabet{}

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}

func TestAlphabet_Less_OtherChars(t *testing.T) {
	lesser := "-abc"
	bigger := "abc"

	comp := Alphabet{}

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}
