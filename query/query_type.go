package query

import "bytes"

// QueryType for druid query
type QueryType int

const (
	// GROUPBY for groupBy query
	GROUPBY QueryType = iota
	// SCAN for scan query
	SCAN
	// SEARCH for search query
	SEARCH
	// TOPN for topN query
	TOPN
)

var queryTypeStrings = []string{
	"groupBy", "scan", "search", "topN",
}

func (q QueryType) name() string {
	return queryTypeStrings[q]
}

func (q QueryType) ordinal() int {
	return int(q)
}

func (q QueryType) values() *[]string {
	return &queryTypeStrings
}

func (q QueryType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(q.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
