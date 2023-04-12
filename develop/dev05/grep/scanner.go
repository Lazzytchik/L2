package grep

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	Options
	Matcher
	Buffer  []string
	cc      int // context counter (how many lines left to print after the match)
	lc      int //	line counter
	OnMatch func(line int, s string)
}

func NewScanner(options Options, matcher Matcher, handler func(line int, s string)) *Scanner {
	scanner := &Scanner{
		Options: options,
		Matcher: matcher,
		OnMatch: handler,
	}

	scanner.cc = 0
	scanner.Buffer = make([]string, 0, options.BeforePrint)

	return scanner
}

func (s *Scanner) Scan(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		s.lc++

		line := scanner.Text()
		target := s.Target

		if s.IgnoreCase {
			line = strings.ToLower(line)
			target = strings.ToLower(target)
		}

		if s.Match(line, target) != s.Invert {
			s.HandleMatch(scanner.Text())
		} else if s.IsAfterMatch() {
			s.HandleAfterMatched(scanner.Text())
		} else {
			s.storeUnmatched(scanner.Text())
		}
	}
}

func (s *Scanner) HandleMatch(line string) {
	s.HandleBeforeMatch()
	s.refreshCounter()
	s.OnMatch(s.lc, line)
}

func (s *Scanner) refreshCounter() {
	s.cc = s.AfterPrint
}

func (s *Scanner) IsAfterMatch() bool {
	return s.cc > 0
}

func (s *Scanner) HandleAfterMatched(line string) {
	s.OnMatch(s.lc, line)
	s.cc--
}

func (s *Scanner) storeUnmatched(line string) {
	if s.BeforePrint == 0 {
		return
	}

	if len(s.Buffer) == s.AfterPrint {
		s.Buffer = s.Buffer[1:]
	}

	s.Buffer = append(s.Buffer, line)
}

func (s *Scanner) HandleBeforeMatch() {

	for _, line := range s.Buffer {
		s.OnMatch(s.lc, line)
	}
	if len(s.Buffer) > 0 {
		s.Buffer = s.Buffer[:1]
	}
}
