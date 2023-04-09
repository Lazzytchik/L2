package dev03

import (
	"log"
)

type Sorter[T comparable] interface {
	Sort([]T) ([]T, error)
}

type StraightSorter struct {
	Comparer  Comparer[string]
	ErrLogger log.Logger
}

func (sorter StraightSorter) Sort(array []string) ([]string, error) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if sorter.Comparer.Less(array[j+1], array[j]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array, nil
}
