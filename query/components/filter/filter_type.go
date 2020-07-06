package filter

import "bytes"

type FilterType int

const (
	SELECTOR FilterType = iota
	COLUMN_COMPARISON
	REGEX
	AND
	OR
	NOT
	JAVASCRIPT
	EXTRACTION
	SEARCH
	IN
	LIKE
	BOUND
	INTERVAL
	TRUE
)

var filterTypeStrings = []string{
	"selector", "columnComparison", "regex", "and", "or", "not", "javascript", "extraction", "search", "in", "like", "bound", "interval", "true",
}

func (f FilterType) name() string {
	return filterTypeStrings[f]
}

func (f FilterType) ordinal() int {
	return int(f)
}

func (f FilterType) values() *[]string {
	return &filterTypeStrings
}

func (f FilterType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(f.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
