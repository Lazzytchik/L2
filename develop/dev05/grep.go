package main

import (
	"errors"
	"flag"
	"lazzytchik/l2/develop/dev05/grep"
	"log"
	"os"
)

/*
	5. 	Утилита grep
		Реализовать утилиту фильтрации по аналогии с консольной

*/

func main() {

	elg := log.New(os.Stderr, "", log.Llongfile)

	options := ParseFlags()
	counter := 0

	scanner := grep.NewScanner(options, ChooseMatcher(options), ChoosePrinter(options, &counter))

	file, err := os.Open(options.FileName)
	if err != nil {
		elg.Printf("grep: error: %s", errors.New("can't open file"))
	}
	defer file.Close()

	scanner.Scan(file)

	PrintCounter(options, counter)

}

func ParseFlags() grep.Options {
	A := flag.Int("A", 0, "print +N strings after match")
	B := flag.Int("B", 0, "print +N strings before match")
	C := flag.Int("C", 0, "print +N strings before and after match")
	c := flag.Bool("c", false, "count match strings")
	i := flag.Bool("i", false, "ignore register")
	v := flag.Bool("v", false, "ignore matches")
	F := flag.Bool("F", false, "exact match")
	n := flag.Bool("n", false, "print line num")

	flag.Parse()

	options := grep.Options{
		AfterPrint:  *C,
		BeforePrint: *C,
		Count:       *c,
		IgnoreCase:  *i,
		Invert:      *v,
		Fixed:       *F,
		LineNum:     *n,
		Target:      flag.Arg(0),
		FileName:    flag.Arg(1),
	}

	if *A > 0 {
		options.AfterPrint = *A
	}

	if *B > 0 {
		options.BeforePrint = *B
	}

	return options
}

func ChooseMatcher(options grep.Options) grep.Matcher {
	if options.Fixed {
		return grep.Exact{}
	}

	return grep.Have{}
}

func ChoosePrinter(options grep.Options, counter *int) func(line int, s string) {
	if options.Count {
		return MakeCounter(counter)
	}

	if options.LineNum {
		return func(line int, s string) {
			log.Printf("line %d: %s", line, s)
		}
	}

	return func(line int, s string) {
		log.Println(s)
	}
}

func MakeCounter(counter *int) func(line int, s string) {
	return func(line int, s string) {
		*counter += 1
	}
}

func PrintCounter(options grep.Options, counter int) {
	if options.Count {
		log.Println(counter)
	}
}
