package aggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"testing"
)

func TestJavascriptAggregatorSerialization(t *testing.T) {
	aggregator := &JavascriptAggregation{
		AggregationType: Javascript,
		Name:            "sum(log(x)*y) + 10",
		FieldNames:       []string{"x", "y"},
		AggregationFn:   "function(current, a, b)      { return current + (Math.log(a) * b); }",
		CombineFn:       "function(partialA, partialB) { return partialA + partialB; }",
		ResetFn:         "function()                   { return 10; }",
	}
	jsonBytes, _ := json.Marshal(aggregator)
	expected := `{
	  "type": "javascript",
	  "name": "sum(log(x)*y) + 10",
	  "fieldNames": ["x", "y"],
	  "fnAggregate" : "function(current, a, b)      { return current + (Math.log(a) * b); }",
	  "fnCombine"   : "function(partialA, partialB) { return partialA + partialB; }",
	  "fnReset"     : "function()                   { return 10; }"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestFilteredAggregationSerialization(t *testing.T) {
	aggregator := &FilteredAggregation{
		AggregationType: Filtered,
		Aggregator:      &Aggregation{
			AggregationType: StringLast,
			Name:            "last_x",
			FieldName:       "x",
			MaxStringBytes:  1024,
		},
		Filter:        &filter.SelectorFilter{
			FilterType: filter.Selector,
			Dimension:  "x",
			Value:      "xx",
		},
	}
	jsonBytes, _ := json.Marshal(aggregator)
	expected := `{
        "type" : "filtered",
        "filter" : {
            "type" : "selector",
			"dimension" : "x",
			"value" : "xx"
        },
        "aggregator" : {
			"type": "stringLast",
			"name": "last_x",
       		"fieldName": "x",
			"maxStringBytes": 1024
        }
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

