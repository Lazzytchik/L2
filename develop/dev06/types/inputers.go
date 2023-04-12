package cut

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
)

type Inputer interface {
	ReadString(delim byte) (string, error)
}

func NewStandard(options Options, logger *log.Logger, oncut func(aftercut string), ctx context.Context) *Standard {
	cutter := Field{
		Index:     options.Fields,
		Delimiter: options.Delimiter,
	}

	return &Standard{
		Reader:  bufio.NewReader(os.Stdin),
		Context: ctx,
		Options: options,
		Logger:  logger,
		Cutter:  cutter,
		OnCut:   oncut,
	}
}

type Standard struct {
	Options
	Reader  *bufio.Reader
	Context context.Context
	Logger  *log.Logger
	Cutter
	OnCut func(aftercut string)
}

func (s *Standard) ReadString() (string, error) {
	text, err := s.Reader.ReadString('\n')
	if err != nil {
		return "", errors.New("input error")
	}

	return text, nil
}

func (s *Standard) Init(signals chan os.Signal) {
	c, cancel := context.WithCancel(s.Context)
	s.Context = c

	go s.handleShutdown(signals, cancel)

	for {
		line, err := s.ReadString()
		if err != nil {
			s.Logger.Println(err)
		} else {
			s.CutTheLine(line)
		}
	}
}

func (s *Standard) CutTheLine(line string) {
	aftercut, sep := s.Cut(line)

	if s.Separated && !sep {
		return
	}

	s.OnCut(aftercut)
}

func (s *Standard) handleShutdown(signals chan os.Signal, cancel context.CancelFunc) {
	select {
	case <-signals:
		cancel()
		os.Exit(0)
	}
}
