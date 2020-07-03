package dimension

type Dimension interface {
	Type() string
}

type SimpleDimension string

func (s SimpleDimension) Type() string {
	return "simple"
}
