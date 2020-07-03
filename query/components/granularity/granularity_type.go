package granularity

import "bytes"

type GranularityType int

const (
	DURATION GranularityType = iota
	PERIOD
)

var granularityTypeStrings = []string{"duration", "period"}

func (g GranularityType) name() string {
	return granularityTypeStrings[g]
}

func (g GranularityType) ordinal() int {
	return int(g)
}

func (g GranularityType) values() *[]string {
	return &granularityTypeStrings
}

func (g GranularityType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(g.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
