package dev03

import (
	"log"
	"strings"
)

type Sorter[T comparable] interface {
	Sort([]T) ([]T, error)
}

type StraightSorter struct {
	Comparer  Comparer[string]
	Options   SortOptions
	ErrLogger *log.Logger
}

func (sorter *StraightSorter) Sort(array []string) ([]string, error) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {

			curr := array[j]
			next := array[j+1]

			if sorter.Options.Trim {
				curr = strings.TrimLeft(curr, " ")
				next = strings.TrimLeft(next, " ")
			}

			if sorter.Options.Column > 1 {
				curr = sorter.extractColumn(curr)
				next = sorter.extractColumn(next)
			}

			if !sorter.Options.Reverse == sorter.Comparer.Less(next, curr) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array, nil
}

func (sorter *StraightSorter) InitComparer() {
	if sorter.Options.Numeric {
		sorter.Comparer = &Numeric{}
	} else if sorter.Options.SuffixNumeric {
		bytes := &Bytes{}
		bytes.SetBaseSuffixes()
		sorter.Comparer = bytes
	} else if sorter.Options.Month {
		month := &Month{}
		month.Init()
		sorter.Comparer = month
	} else {
		sorter.Comparer = Alphabet{}
	}
}

func (sorter *StraightSorter) extractColumn(s string) string {
	fields := strings.Fields(s)
	if len(fields) < sorter.Options.Column {
		return ""
	}

	return fields[sorter.Options.Column-1]
}
