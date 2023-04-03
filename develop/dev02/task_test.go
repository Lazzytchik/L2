package task02

import (
	"testing"
)

func TestUnpackGeneralPositiveCheck(t *testing.T) {
	example := "a4bb8c7aa"
	must := "aaaabbbbbbbbbcccccccaa"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}

func TestString_Unpack_UnicodePositiveCheck(t *testing.T) {
	example := "a4b假8c7aa"
	must := "aaaab假假假假假假假假cccccccaa"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}

func TestString_Unpack_IncorrectInputNegativeCheck(t *testing.T) {
	example := "35"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); err == nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", "error", result)
	}
}

func TestString_Unpack_EmptyInputPositiveCheck(t *testing.T) {
	example := ""
	must := ""

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}
