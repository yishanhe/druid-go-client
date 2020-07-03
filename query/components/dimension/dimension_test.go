package dimension

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultDimension(t *testing.T) {
	var dimSpec Dimension
	dimSpec = &DefaultDimension{
		DimensionType: DEFAULT,
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

func TestDefaultDimensionOutputType(t *testing.T) {
	var dimSpec Dimension
	dimSpec = &DefaultDimension{
		DimensionType: REGEX_FILTERED,
		Dimension:     "field1",
		OutputName:    "field1Out",
		OutputType:    LONG,
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
