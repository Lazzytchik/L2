package cut

type Options struct {
	Fields    int
	Delimiter string
	Separated bool
}

func (o Options) Fielded() bool {
	return o.Fields > 0
}
