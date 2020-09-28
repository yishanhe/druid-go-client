package limit

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/sorting"
	"testing"
)

func TestLimitSpecSerialization(t *testing.T) {
	limit := &DefaultLimit{
		LimitType: Default,
		Limit:     10000,
		Columns: []sorting.OrderByColumn{
			{
				Dimension:      "dim1",
				Direction:      sorting.Ascending,
				DimensionOrder: sorting.Lexicographic,
			},
		},
	}
	jsonBytes, _ := json.Marshal(limit)
	expected := `{
		"type"    : "default",
		"limit"   : 10000,
		"columns" : [{
			"dimension" : "dim1",
			"direction" : "ascending",
			"dimensionOrder" : "lexicographic"
	}]
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}
