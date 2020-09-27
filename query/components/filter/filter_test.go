package filter

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/extraction"
	"github.com/yishanhe/druid-go-client/query/components/lookup"
	"github.com/yishanhe/druid-go-client/query/components/search"
)

func TestExtractionFilterSerialization(t *testing.T) {
	var filter Filter
	filter = &ExtractionFilter{
		FilterType: Extraction,
		Dimension:  "product",
		Value:      "bar_1",
		Extraction: &extraction.LookupExtraction{
			ExtractionType: extraction.Lookup,
			Lookup: &lookup.MapLookup{
				LookupType: lookup.Map,
				Map: map[string]string{
					"product_1": "bar_1",
					"product_5": "bar_1",
					"product_3": "bar_1",
				},
			},
		},
	}

	jsonBytes, _ := json.Marshal(filter)

	expected := `{
			"type": "extraction",
			"dimension": "product",
			"value": "bar_1",
			"extractionFn": {
				"type": "lookup",
				"lookup": {
					"type": "map",
					"map": {
						"product_1": "bar_1",
						"product_5": "bar_1",
						"product_3": "bar_1"
					}
				}
			}
		}`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestJavascriptFilterSerialization(t *testing.T) {
	var filter Filter
	filter = &JavaScriptFilter{
		FilterType: Javascript,
		Dimension:  "name",
		Function:   "function(x) { return(x >= 'bar' && x <= 'foo') }",
	}

	jsonBytes, _ := json.Marshal(filter)

	expected := `{
		"type" : "javascript",
		"dimension" : "name",
		"function" : "function(x) { return(x >= 'bar' && x <= 'foo') }"
	}`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestIntervalFilterSerialization(t *testing.T) {
	var filter Filter
	filter = &IntervalFilter{
		FilterType: Interval,
		Dimension:  "__time",
		Intervals: []string{"2014-10-01T00:00:00.000Z/2014-10-07T00:00:00.000Z",
			"2014-11-15T00:00:00.000Z/2014-11-16T00:00:00.000Z"},
	}

	jsonBytes, _ := json.Marshal(filter)

	expected := `{
		"type" : "interval",
		"dimension" : "__time",
		"intervals" : [
		  "2014-10-01T00:00:00.000Z/2014-10-07T00:00:00.000Z",
		  "2014-11-15T00:00:00.000Z/2014-11-16T00:00:00.000Z"
		]
	}`

	require.JSONEq(t, expected, string(jsonBytes))

}
func TestSearchFilterSerialization(t *testing.T) {
	filter := &SearchFilter{
		FilterType: Search,
		Dimension:  "product",
		Query: &search.InsensitiveContainsSearch{
			SearchType: search.InsensitiveContains,
			Value:      "foo",
		},
	}
	jsonBytes, _ := json.Marshal(filter)
	expected := `{
        "type": "search",
        "dimension": "product",
        "query": {
          "type": "insensitive_contains",
          "value": "foo"
        }
    }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestInFilterSerialization(t *testing.T) {
	filter := &InFilter{
		FilterType: In,
		Dimension:  "outlaw",
		Values:     []string{"Good", "Bad", "Ugly"},
	}
	jsonBytes, _ := json.Marshal(filter)
	expected := `{
		"type": "in",
		"dimension": "outlaw",
		"values": ["Good", "Bad", "Ugly"]
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestLikeFilterSerialization(t *testing.T) {
	filter := &LikeFilter{
		FilterType: Like,
		Dimension:  "last_name",
		Pattern:    "D%",
	}
	jsonBytes, _ := json.Marshal(filter)
	expected := `{
          "type": "like",
          "dimension": "last_name",
          "pattern": "D%"
        }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestTrueFilterSerialization(t *testing.T) {
	filter := &TrueFilter{
		FilterType: True,
	}
	jsonBytes, _ := json.Marshal(filter)
	expected := `{"type": "true"}`
	require.JSONEq(t, expected, string(jsonBytes))
}
