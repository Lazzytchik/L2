package main

import (
	"bufio"
	"errors"
	"flag"
	file "lazzytchik/l2/common/files"
	dev03 "lazzytchik/l2/develop/dev03/types"
	"log"
	"os"
)

func main() {
	from := "data.txt"

	//	Init logger
	errLogger := log.New(os.Stderr, "sort: ", log.Ldate)

	//	Init sort options
	options := ParseFlagOptions()
	if options.IsIncompatible() {
		errLogger.Printf("options \"%s\" incompatible", options.String())
		return
	}

	//	Get data from file
	raw, err := file.Lines(from)
	if err != nil {
		errLogger.Fatalf(err.Error())
	}
	for _, s := range raw {
		log.Println(s)
	}

	//	Sort extracted data
	sorter := dev03.StraightSorter{
		Options:   options,
		ErrLogger: errLogger,
	}
	sorter.InitComparer()

	sorted, _ := sorter.Sort(raw)

	//	Write it to the new file
	log.Println("____SORTED____")
	for _, s := range sorted {
		log.Println(s)
	}
}

func getDataFromFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("can't open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ParseFlagOptions() dev03.SortOptions {
	column := flag.Int("k", 1, "number of word to sort by")
	numeric := flag.Bool("n", false, "numeric sort")
	reverse := flag.Bool("r", false, "reverse sort")
	omitDuplicates := flag.Bool("u", false, "omit duplicates")
	month := flag.Bool("M", false, "month sort")
	trim := flag.Bool("b", false, "trim left spaces")
	checkSorted := flag.Bool("c", false, "check if sorted")
	suffixNumeric := flag.Bool("h", false, "suffixes supported numeric sort")

	flag.Parse()

	options := dev03.SortOptions{
		Column:         *column,
		Numeric:        *numeric,
		Reverse:        *reverse,
		OmitDuplicates: *omitDuplicates,
		Month:          *month,
		Trim:           *trim,
		CheckSorted:    *checkSorted,
		SuffixNumeric:  *suffixNumeric,
	}

	return options
}
