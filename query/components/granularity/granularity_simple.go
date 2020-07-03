package granularity

import "bytes"

type SimpleGranularity int

const (
	ALL SimpleGranularity = iota
	NONE
	SECOND
	MINUTE
	FIFTEEN_MINUTE
	THIRTY_MINUTE
	HOUR
	DAY
	WEEK
	MONTH
	QUARTER
	YEAR
)

var simpleGranularityStrings = []string{
	"all", "none", "second", "minute", "fifteen_minute", "thirty_minute", "hour", "day", "week", "month", "quarter", "year",
}

func (s SimpleGranularity) name() string {
	return simpleGranularityStrings[s]
}

func (s SimpleGranularity) ordinal() int {
	return int(s)
}

func (s SimpleGranularity) values() *[]string {
	return &simpleGranularityStrings
}

func (s SimpleGranularity) Type() string {
	return "SimpleGranularity"
}

func (s SimpleGranularity) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
