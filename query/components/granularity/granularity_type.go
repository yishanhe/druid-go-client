package granularity

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type GranularityType int

const (
	Duration GranularityType = iota
	Period
)

var granularityTypeStrings = []string{"duration", "period"}

func (g GranularityType) Name() string {
	return granularityTypeStrings[g]
}

func (g GranularityType) Ordinal() int {
	return int(g)
}

func (g GranularityType) Values() *[]string {
	return &granularityTypeStrings
}

func (g GranularityType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(g.Name())
}
