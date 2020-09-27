package dimension

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

// DimensionType for druid query
type DimensionType int

const (
	Default DimensionType = iota
	Extraction
	Lookup
	ListFiltered
	RegexFiltered
	PrefixFiltered
)

var dimensionTypeStrings = []string{"default", "extraction", "lookup", "listFiltered", "regexFiltered", "prefixFiltered"}

func (d DimensionType) Name() string {
	return dimensionTypeStrings[d]
}

func (d DimensionType) Ordinal() int {
	return int(d)
}

func (d DimensionType) Values() *[]string {
	return &dimensionTypeStrings
}

func (d DimensionType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(d.Name())
}
