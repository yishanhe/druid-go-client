package extraction

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type ExtractionType int

const (
	Regex ExtractionType = iota
	Partial
	SearchQuery
	Substring
	Strlen
	TimeFormat
	Time
	Javascript
	RegisteredLookup
	Lookup
	Cascade
	StringFormat
	Upper
	Lower
	Bucket
)

var extractionTypeStrings = []string{
	"regex", "partial", "searchQuery", "substring", "strlen", "timeFormat", "time", "javascript", "registeredLookup", "lookup", "cascade", "stringFormat",
	"upper", "lower", "bucket",
}

func (e ExtractionType) Name() string {
	return extractionTypeStrings[e]
}

func (e ExtractionType) Ordinal() int {
	return int(e)
}

func (e ExtractionType) Values() *[]string {
	return &extractionTypeStrings
}

func (e ExtractionType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON((e.Name()))
}
