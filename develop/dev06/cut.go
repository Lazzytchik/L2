package main

import (
	"context"
	"flag"
	cut "lazzytchik/l2/develop/dev06/types"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
	6. 	Реализовать утилиту аналог консольной команды cut (man cut).
		Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

		Реализовать поддержку утилитой следующих ключей:
			-f - "fields" - выбрать поля (колонки)
			-d - "delimiter" - использовать другой разделитель
			-s - "separated" - только строки с разделителем

*/

func main() {
	options := ParseFlags()

	elg := log.New(os.Stderr, "cut: ", log.Ldate)

	if !options.Fielded() {
		elg.Fatalf("you should choose cut method like -f")
	}

	signals := make(chan os.Signal, 1)
	CatchSignals(signals)

	inputer := cut.NewStandard(options, elg, func(aftercut string) {
		log.Println(aftercut)
	}, context.Background())

	inputer.Init(signals)

}

func ParseFlags() cut.Options {
	f := flag.Int("f", 0, "cut exact field")
	d := flag.String("d", "\t", "set delimiter")
	s := flag.Bool("s", false, "cut only lines containing delimiter")

	flag.Parse()

	options := cut.Options{
		Fields:    *f,
		Delimiter: *d,
		Separated: *s,
	}

	return options
}

func CatchSignals(signals chan os.Signal) {
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGTERM)
}
