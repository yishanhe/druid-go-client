package virtualcolumn

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/common"
	"testing"
)

func TestExpressionVirtualColumnSerialization(t *testing.T) {
	vc := &ExpressionVirtualColumn{
		VirtualColumnType: Expression,
		Name:              "fooPage",
		Expression:        "concat('foo' + page)",
		OutputType:        common.String,
	}
	jsonBytes, _ := json.Marshal(vc)
	expected := `    {
      "type": "expression",
      "name": "fooPage",
      "expression": "concat('foo' + page)",
      "outputType": "STRING"
    }`
	require.JSONEq(t, expected, string(jsonBytes))
}
