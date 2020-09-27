package sorting

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type SortingOrder int

const (
	Lexicographic SortingOrder = iota
	Alphanumeric
	Numeric
	Strlen
	Version
)

var sortingOrderStrings = []string{
	"lexicographic", "alphanumeric", "numeric", "strlen", "version",
}

func (s SortingOrder) Name() string {
	return sortingOrderStrings[s]
}

func (s SortingOrder) Ordinal() int {
	return int(s)
}

func (s SortingOrder) Values() *[]string {
	return &sortingOrderStrings
}

func (s SortingOrder) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(s.Name())
}
