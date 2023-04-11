package grep

type Options struct {
	AfterPrint  int
	BeforePrint int
	Count       bool
	IgnoreCase  bool
	Invert      bool
	Fixed       bool
	LineNum     bool
	Target      string
	FileName    string
}
