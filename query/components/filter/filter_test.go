package filter

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntervalFilter(t *testing.T) {
	var filter Filter
	filter = &IntervalFilter{
		FilterType: INTERVAL,
		Dimension:  "__time",
		Intervals: []string{"2014-10-01T00:00:00.000Z/2014-10-07T00:00:00.000Z",
			"2014-11-15T00:00:00.000Z/2014-11-16T00:00:00.000Z"},
	}

	expected := `{
		"type" : "interval",
		"dimension" : "__time",
		"intervals" : [
		  "2014-10-01T00:00:00.000Z/2014-10-07T00:00:00.000Z",
		  "2014-11-15T00:00:00.000Z/2014-11-16T00:00:00.000Z"
		]
	}`

	jsonBytes, _ := json.Marshal(filter)

	require.JSONEq(t, expected, string(jsonBytes))

}
