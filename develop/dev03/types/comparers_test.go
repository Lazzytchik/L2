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

func TestNumeric_Less(t *testing.T) {
	lesser := "82abc"
	bigger := "382abc"

	comp := Numeric{}

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}

func TestNumeric_Less_Same(t *testing.T) {
	lesser := "8927"
	bigger := "8927"

	comp := Numeric{}

	if comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}

func TestMonth_Less(t *testing.T) {
	lesser := "apr"
	bigger := "Sep"

	comp := Month{
		Alphabet: Alphabet{},
	}
	comp.SetEngMonths()

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}

func TestMonth_ExtractMonth(t *testing.T) {
	comp := Month{}
	comp.Init()

	want := 6
	input := "Jun"

	if got := comp.ExtractMonth(input); got != want {
		t.Fatalf("Wrong answer: want %d, got %d", want, got)
	}
}

func TestMonth_ExtractMonth_Negative(t *testing.T) {
	comp := Month{}
	comp.Init()

	want := 4
	input := "Jun"

	if got := comp.ExtractMonth(input); got == want {
		t.Fatalf("Wrong answer: want %d, got %d", want, got)
	}
}

func TestBytes_Less(t *testing.T) {
	lesser := "2100"
	bigger := "20K"

	comp := Bytes{}
	comp.SetBaseSuffixes()

	if !comp.Less(lesser, bigger) {
		t.Fatalf("wrong answer: %s must be higher than %s", lesser, bigger)
	}
}
