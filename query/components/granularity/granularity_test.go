package granularity

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDurationGranularitySerialization(t *testing.T) {
	gran := &DurationGranularity{
		GranularityType: Duration,
		Duration:        7200000,
	}
	jsonBytes, _ := json.Marshal(gran)
	expected := `{"type": "duration", "duration": 7200000}`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestPeriodGranularitySerialization(t *testing.T) {
	gran := &PeriodGranularity{
		GranularityType: Period,
		Period:          "P3M",
		TimeZone:        "America/Los_Angeles",
		Origin:          "2012-02-01T00:00:00-08:00",
	}
	jsonBytes, _ := json.Marshal(gran)
	expected := `{"type": "period", "period": "P3M", "timeZone": "America/Los_Angeles", "origin": "2012-02-01T00:00:00-08:00"}`
	require.JSONEq(t, expected, string(jsonBytes))
}
