package dimension

type DefaultDimension struct {
	DimensionType DimensionType `json:"type"`
	Dimension     string        `json:"dimension"`
	OutputName    string        `json:"outputName"`
	OutputType    OutputType    `json:"outputType,omitempty"`
}

func (d *DefaultDimension) Type() string {
	return DEFAULT.name()
}
