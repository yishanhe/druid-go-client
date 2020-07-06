package sorting

import "bytes"

type SortingOrder int

const (
	LEXICOGRAPHIC SortingOrder = iota
	ALPHA_NUMERIC
	NUMERIC
	STRLEN
	VERSION
)

var sortingOrderStrings = []string{
	"lexicographic", "alphanumeric", "numeric", "strlen", "version",
}

func (s SortingOrder) name() string {
	return sortingOrderStrings[s]
}

func (s SortingOrder) ordinal() int {
	return int(s)
}

func (s SortingOrder) values() *[]string {
	return &sortingOrderStrings
}

func (s SortingOrder) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
