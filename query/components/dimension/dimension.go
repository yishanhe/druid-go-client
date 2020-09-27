package dimension

type Dimension interface {
	Type() string
}

type SimpleDimension string

func (s SimpleDimension) Type() string {
	return "simple"
}

type DefaultDimension struct {
	DimensionType DimensionType `json:"type"`
	Dimension     string        `json:"dimension"`
	OutputName    string        `json:"outputName"`
	OutputType    OutputType    `json:"outputType,omitempty"`
}

func (d DefaultDimension) Type() string {
	return Default.Name()
}
