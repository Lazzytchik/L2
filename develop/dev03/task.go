package main

import (
	"bufio"
	"errors"
	"flag"
	dev03 "lazzytchik/l2/develop/dev03/types"
	"log"
	"os"
)

func Sort(from string, outputFile string) (string, error) {
	//	Init logger
	errLogger := log.New(os.Stderr, "sort: ", log.Ldate)

	//	Init sort options
	options := dev03.SortOptions{}
	options.Init()
	flag.Parse()

	//	Get data from file
	_, err := getDataFromFile(from)
	if err != nil {
		errLogger.Fatalf(err.Error())
	}

	//	Sort extracted data

	//	Write it to the new file
	return outputFile, nil
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
