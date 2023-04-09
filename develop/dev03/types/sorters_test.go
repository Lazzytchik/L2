package dev03

import "testing"

func TestStraightSorter_Sort(t *testing.T) {
	want := []string{
		"8bc",
		"abc",
		"Abc",
	}

	sorter := StraightSorter{
		Comparer: Alphabet{},
	}

	have, err := sorter.Sort([]string{
		"abc",
		"Abc",
		"8bc",
	})

	if err != nil {
		t.Fatalf("error: %s. want: %s, got %s", err, want, err)
	}

	if !testEq(have, want) {
		t.Fatalf("wrong result: want %s, got %s", want, have)
	}

}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
