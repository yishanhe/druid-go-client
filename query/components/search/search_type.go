package search

import "bytes"

type SearchType int

const (
	INSENSITIVE_CONTAINS SearchType = iota
	FRAGMENT
	CONTAINS
	REGEX
)

var searchTypeStrings = []string{
	"insensitive_contains", "fragment", "contains", "regex",
}

func (s SearchType) name() string {
	return searchTypeStrings[s]
}

func (s SearchType) ordinal() int {
	return int(s)
}

func (s SearchType) values() *[]string {
	return &searchTypeStrings
}

func (s SearchType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
