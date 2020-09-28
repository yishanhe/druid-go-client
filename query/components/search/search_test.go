package search

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFragmentSearchSerialization(t *testing.T) {
	search := &FragmentSearch{
		SearchType:    Fragment,
		CaseSensitive: true,
		Values:        []string{"fragment1", "fragment2"},
	}
	jsonBytes, _ := json.Marshal(search)
	expected := `{
	  "type" : "fragment",
	  "case_sensitive" : true,
	  "values" : ["fragment1", "fragment2"]
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestContainsSearchSerialization(t *testing.T) {
	search := &ContainsSearch{
		SearchType:    Contains,
		CaseSensitive: true,
		Value:         "some_value",
	}
	jsonBytes, _ := json.Marshal(search)
	expected := `{
	  "type"  : "contains",
	  "case_sensitive" : true,
	  "value" : "some_value"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestRegexSearchSerialization(t *testing.T) {
	search := &RegexSearch{
		SearchType: Regex,
		Pattern:    "some_pattern",
	}
	jsonBytes, _ := json.Marshal(search)
	expected := `{
	  "type"  : "regex",
	  "pattern" : "some_pattern"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}
