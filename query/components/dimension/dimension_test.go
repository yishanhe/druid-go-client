package dimension

import (
	"encoding/json"
	"github.com/yishanhe/druid-go-client/query/components/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/lookup"
)

func TestDefaultDimensionSerialization(t *testing.T) {
	var dimSpec Dimension
	dimSpec = &DefaultDimension{
		DimensionType: Default,
		Dimension:     "field1",
		OutputName:    "field1Out",
	}
	expected := `{
		"type": "default",
		"dimension": "field1",
		"outputName": "field1Out"
	}`
	jsonBytes, _ := json.Marshal(dimSpec)
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestDefaultDimensionOutputTypeSerialization(t *testing.T) {
	var dimSpec Dimension
	dimSpec = &DefaultDimension{
		DimensionType: RegexFiltered,
		Dimension:     "field1",
		OutputName:    "field1Out",
		OutputType:    common.Long,
	}
	expected := `{
		"type": "regexFiltered",
		"dimension": "field1",
		"outputName": "field1Out",
		"outputType": "LONG"
	}`
	jsonBytes, err := json.Marshal(dimSpec)
	if err != nil {
		assert.Fail(t, "Failed")
	}
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestLookupDimensionSerialization(t *testing.T) {
	dim := &LookupDimension{
		DimensionType:           Lookup,
		Dimension:               "dimensionName",
		OutputName:              "dimensionOutputName",
		ReplaceMissingValueWith: "missing_value",
		RetainMissingValue:      true,
		Lookup: &lookup.MapLookup{
			LookupType: lookup.Map,
			Map: map[string]string{
				"key": "value",
			},
			IsOneToOne: true,
		},
	}
	jsonBytes, err := json.Marshal(dim)
	if err != nil {
		assert.Fail(t, "Serialization Failed")
	}
	expected := `{
		"type":"lookup",
		"dimension":"dimensionName",
		"outputName":"dimensionOutputName",
		"replaceMissingValueWith":"missing_value",
		"retainMissingValue":true,
		"lookup":{"type": "map", "map":{"key":"value"}, "isOneToOne":true}
	  }`
	require.JSONEq(t, expected, string(jsonBytes))
}
