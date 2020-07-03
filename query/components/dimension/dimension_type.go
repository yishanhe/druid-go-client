package dimension

import "bytes"

// DimensionType for druid query
type DimensionType int

const (
	DEFAULT DimensionType = iota
	EXTRACTION
	LOOKUP
	LIST_FILTERED
	REGEX_FILTERED
	PREFIX_FILTERED
)

var dimensionTypeStrings = []string{"default", "extraction", "lookup", "listFiltered", "regexFiltered", "prefixFiltered"}

func (d DimensionType) name() string {
	return dimensionTypeStrings[d]
}

func (d DimensionType) ordinal() int {
	return int(d)
}

func (d DimensionType) values() *[]string {
	return &dimensionTypeStrings
}

func (d DimensionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(d.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
