package extraction

import "bytes"

type ExtractionType int

const (
	REGEX ExtractionType = iota
	PARTIAL
	SEARCH_QUERY
	SUBSTRING
	STRLEN
	TIME_FORMAT
	TIME
	JAVASCRIPT
	REGISTERED_LOOKUP
	LOOKUP
	CASCADE
	STRING_FORMAT
	UPPER
	LOWER
	BUCKET
)

var extractionTypeStrings = []string{
	"regex", "partial", "searchQuery", "substring", "strlen", "timeFormat", "time", "javascript", "registeredLookup", "lookup", "cascade", "stringFormat",
	"upper", "lower", "bucket",
}

func (e ExtractionType) name() string {
	return extractionTypeStrings[e]
}

func (e ExtractionType) ordinal() int {
	return int(e)
}

func (e ExtractionType) values() *[]string {
	return &extractionTypeStrings
}

func (e ExtractionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(e.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
