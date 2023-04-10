package dev03

import (
	"flag"
	"fmt"
	"strings"
)

type SortOptions struct {
	Column         int
	Numeric        bool
	Reverse        bool
	OmitDuplicates bool
	Month          bool
	Trim           bool
	CheckSorted    bool
	SuffixNumeric  bool
}

func (o *SortOptions) Parse() {
	flag.Parse()
}

func (o *SortOptions) IsIncompatible() bool {
	return (o.Month && o.SuffixNumeric) || (o.Numeric && o.SuffixNumeric) || (o.Month && o.Numeric)
}

func (o *SortOptions) String() string {
	var result strings.Builder

	if o.Column > 1 {
		result.WriteString(fmt.Sprintf("-k %d ", o.Column))
	}

	if o.Numeric {
		result.WriteString("-n ")
	}

	if o.Trim {
		result.WriteString("-b ")
	}

	if o.Month {
		result.WriteString("-M ")
	}

	if o.CheckSorted {
		result.WriteString("-c ")
	}

	if o.SuffixNumeric {
		result.WriteString("-h ")
	}

	if o.Reverse {
		result.WriteString("-r ")
	}

	if o.OmitDuplicates {
		result.WriteString("-u ")
	}

	return result.String()
}
