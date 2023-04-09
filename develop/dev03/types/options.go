package dev03

import "flag"

type SortOptions struct {
	Column         *int
	Numeric        *bool
	Reverse        *bool
	OmitDuplicates *bool
	Month          *bool
	Trim           *bool
	CheckSorted    *bool
	SuffixNumeric  *bool
}

func (o SortOptions) Init() {
	o.Column = flag.Int("k", 1, "number of word to sort by")
	o.Numeric = flag.Bool("n", false, "numeric sort")
	o.Reverse = flag.Bool("r", false, "reverse sort")
	o.OmitDuplicates = flag.Bool("u", false, "omit duplicates")
	o.Month = flag.Bool("M", false, "month sort")
	o.Trim = flag.Bool("b", false, "trim left spaces")
	o.CheckSorted = flag.Bool("c", false, "check if sorted")
	o.SuffixNumeric = flag.Bool("h", false, "suffixes supported numeric sort")
}
