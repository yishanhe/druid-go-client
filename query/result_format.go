package query

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type ResultFormat int

const (
	List ResultFormat = iota
	CompactedList
)

var resultFormatStrings = []string{"list", "compactedList"}

func (r ResultFormat) Name() string {
	return resultFormatStrings[r]
}

func (r ResultFormat) Ordinal() int {
	return int(r)
}

func (r ResultFormat) Values() *[]string {
	return &resultFormatStrings
}

func (r ResultFormat) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(r.Name())
}
