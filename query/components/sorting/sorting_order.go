package sorting

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type SortingOrder struct {
	SortingOrderType SortingOrderType `json:"type"`
}

type SortingOrderType int

const (
	Lexicographic SortingOrderType = iota
	Alphanumeric
	Numeric
	Strlen
	Version
)

var sortingOrderStrings = []string{
	"lexicographic", "alphanumeric", "numeric", "strlen", "version",
}

func (s SortingOrderType) Name() string {
	return sortingOrderStrings[s]
}

func (s SortingOrderType) Ordinal() int {
	return int(s)
}

func (s SortingOrderType) Values() *[]string {
	return &sortingOrderStrings
}

func (s SortingOrderType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(s.Name())
}
