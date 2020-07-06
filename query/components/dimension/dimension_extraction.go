package dimension

import "github.com/yishanhe/druid-go-client/query/components/extraction"

type ExtractionDimension struct {
	DimensionType DimensionType         `json:"type"`
	Dimension     string                `json:"dimension"`
	OutputName    string                `json:"outputName"`
	OutputType    OutputType            `json:"outputType,omitempty`
	ExtractionFn  extraction.Extraction `json:"extractionFn"`
}

func (d *ExtractionDimension) Type() string {
	return EXTRACTION.name()
}
