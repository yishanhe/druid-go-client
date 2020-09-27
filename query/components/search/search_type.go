package search

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type SearchType int

const (
	InsensitiveContains SearchType = iota
	Fragment
	Contains
	Regex
)

var searchTypeStrings = []string{
	"insensitive_contains", "fragment", "contains", "regex",
}

func (s SearchType) Name() string {
	return searchTypeStrings[s]
}

func (s SearchType) Ordinal() int {
	return int(s)
}

func (s SearchType) Values() *[]string {
	return &searchTypeStrings
}

func (s SearchType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(s.Name())
}
