package common

import "github.com/yishanhe/druid-go-client/pkg/enum"

type Bound int

const (
	MaxTime Bound = iota
	MinTime
)

var boundStrings = []string{"maxTime", "minTime"}

func (b Bound) Name() string {
	return boundStrings[b]
}

func (b Bound) Ordinal() int {
	return int(b)
}

func (b Bound) Values() *[]string {
	return &boundStrings
}

func (b Bound) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(b.Name())
}
