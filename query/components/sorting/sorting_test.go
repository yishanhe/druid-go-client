package sorting

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderByColumnSerialization(t *testing.T) {
	orderby := &OrderByColumn{
		Dimension:      "dim1",
		Direction:      Ascending,
		DimensionOrder: Lexicographic,
	}
	jsonBytes, _ := json.Marshal(orderby)
	expected := `{
		"dimension" : "dim1",
		"direction" : "ascending",
		"dimensionOrder" : "lexicographic"
	}`
	require.JSONEq(t, expected, string(jsonBytes))
}
