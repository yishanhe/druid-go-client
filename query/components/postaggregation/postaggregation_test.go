package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFieldAccessPostAggregationSerialization(t *testing.T) {
	agg := &FieldAccessPostAggregation{
		PostAggregationType: FieldAccess,
		Name:                "outputName",
		FieldName:           "aggregatorName",
	}
	jsonBytes, _ := json.Marshal(agg)
	expected := `{
		"type"  : "fieldAccess",
		"name"  : "outputName",
		"fieldName": "aggregatorName"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestJavascriptPostAggregationSerialization(t *testing.T) {
	agg := &JavascriptPostAggregation{
		PostAggregationType: Javascript,
		Name:                "absPercent",
		FieldNames:          []string{"delta", "total"},
		Function:            "function(delta, total) { return 100 * Math.abs(delta) / total; }",
	}
	jsonBytes, _ := json.Marshal(agg)
	expected := `{
	  "type": "javascript",
	  "name": "absPercent",
	  "fieldNames": ["delta", "total"],
	  "function": "function(delta, total) { return 100 * Math.abs(delta) / total; }"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestGreatestAggregationSerialization(t *testing.T) {
	agg := &GreatestPostAggregation{
		PostAggregationType: DoubleGreatest,
		Name:                "outputName",
		Fields: []PostAggregator{
			&FieldAccessPostAggregation{
				PostAggregationType: FieldAccess,
				Name:                "outputName",
				FieldName:           "aggregatorName",
			},
		},
	}
	jsonBytes, _ := json.Marshal(agg)
	expected := `{
		"type"  : "doubleGreatest",
		"name"  : "outputName",
		"fields": [{
	  		"type"  : "fieldAccess",
	  		"name"  : "outputName",
			"fieldName": "aggregatorName"
			}]
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestArithmeticPostAggregationSerialization(t *testing.T) {
	agg := &ArithmeticPostAggregation{
		PostAggregationType: Arithmetic,
		Name:                "average_users_per_row",
		Function:            "/",
		Fields: []PostAggregator{
			&FieldAccessPostAggregation{
				PostAggregationType: HyperUniqueCardinality,
				FieldName:           "unique_users",
			},
			&FieldAccessPostAggregation{
				PostAggregationType: FieldAccess,
				Name:                "rows",
				FieldName:           "rows",
			},
		},
	}
	jsonBytes, _ := json.Marshal(agg)
	expected := `{
    "type"   : "arithmetic",
    "name"   : "average_users_per_row",
    "fn"     : "/",
    "fields" : [
      { "type" : "hyperUniqueCardinality", "fieldName" : "unique_users" },
      { "type" : "fieldAccess", "name" : "rows", "fieldName" : "rows" }
    ]
  }`
	require.JSONEq(t, expected, string(jsonBytes))
}
