package granularity

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

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

func (s SimpleGranularity) Name() string {
	return simpleGranularityStrings[s]
}

func (s SimpleGranularity) Ordinal() int {
	return int(s)
}

func (s SimpleGranularity) Values() *[]string {
	return &simpleGranularityStrings
}

func (s SimpleGranularity) Type() string {
	return "SimpleGranularity"
}

func (s SimpleGranularity) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(s.Name())
}
