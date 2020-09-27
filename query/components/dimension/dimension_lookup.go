package dimension

import (
	"github.com/yishanhe/druid-go-client/query/components/lookup"
)

type LookupDimension struct {
	DimensionType           DimensionType `json:"type"`
	Dimension               string        `json:"dimension"`
	OutputName              string        `json:"outputName"`
	ReplaceMissingValueWith string        `json:"replaceMissingValueWith"`
	RetainMissingValue      bool          `json:"retainMissingValue"`
	Lookup                  lookup.Lookup `json:"lookup"`
}

func (l LookupDimension) Type() string {
	return Lookup.Name()
}
