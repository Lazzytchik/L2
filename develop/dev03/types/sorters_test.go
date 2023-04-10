package dev03

import "testing"

func TestStraightSorter_Sort(t *testing.T) {
	want := []string{
		"8bc",
		"abc",
		"Abc",
	}

	options := SortOptions{}

	sorter := StraightSorter{
		Comparer: Alphabet{},
		Options:  options,
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

func TestStraightSorter_Sort_Reverse(t *testing.T) {
	want := []string{
		"Abc",
		"abc",
		"8bc",
	}

	options := SortOptions{
		Reverse: true,
	}

	sorter := StraightSorter{
		Comparer: Alphabet{},
		Options:  options,
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

func TestStraightSorter_Sort_Numeric(t *testing.T) {
	want := []string{
		"8",
		"82",
		"187",
	}

	options := SortOptions{
		Numeric: true,
	}

	sorter := StraightSorter{
		Options: options,
	}
	sorter.InitComparer()

	have, err := sorter.Sort([]string{
		"82",
		"187",
		"8",
	})

	if err != nil {
		t.Fatalf("error: %s. want: %s, got %s", err, want, err)
	}

	if !testEq(have, want) {
		t.Fatalf("wrong result: want %s, got %s", want, have)
	}
}

func TestStraightSorter_Sort_Numeric_WithText(t *testing.T) {
	want := []string{
		"8aa",
		"82aaa",
		"187aaaa",
	}

	options := SortOptions{
		Numeric: true,
	}

	sorter := StraightSorter{
		Options: options,
	}
	sorter.InitComparer()

	have, err := sorter.Sort([]string{
		"82aaa",
		"187aaaa",
		"8aa",
	})

	if err != nil {
		t.Fatalf("error: %s. want: %s, got %s", err, want, err)
	}

	if !testEq(have, want) {
		t.Fatalf("wrong result: want %s, got %s", want, have)
	}
}

func TestStraightSorter_Sort_Bytes_Positive(t *testing.T) {
	want := []string{
		"a",
		"8K",
		"1M",
		"20G",
	}

	options := SortOptions{
		SuffixNumeric: true,
	}

	sorter := StraightSorter{
		Options: options,
	}
	sorter.InitComparer()

	have, err := sorter.Sort([]string{
		"20G",
		"a",
		"8K",
		"1M",
	})

	if err != nil {
		t.Fatalf("error: %s. want: %s, got %s", err, want, err)
	}

	if !testEq(have, want) {
		t.Fatalf("wrong result: want %s, got %s", want, have)
	}
}

func TestStraightSorter_Sort_Column(t *testing.T) {
	want := []string{
		"man aac zc",
		"aude Abc abc",
		"zbd ed aaa",
		"dude mbc dude",
	}

	options := SortOptions{
		Column: 2,
	}

	sorter := StraightSorter{
		Options: options,
	}
	sorter.InitComparer()

	have, err := sorter.Sort([]string{
		"dude mbc dude",
		"aude Abc abc",
		"man aac zc",
		"zbd ed aaa",
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
