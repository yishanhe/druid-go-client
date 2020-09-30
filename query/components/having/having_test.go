package having

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"testing"
)

func TestFilteredHavingSpecSerialization(t *testing.T) {
	having := &FilterHavingSpec{
		HavingType: Filter,
		Filter: &filter.SelectorFilter{
			FilterType: filter.Selector,
			Dimension:  "a",
			Value:      "b",
		},
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
            "type" : "filter",
            "filter" : {
              "type": "selector",
              "dimension" : "a",
              "value" : "b"
            }
        }`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestNumericHavingSpecSerialization(t *testing.T) {
	having := &NumericHavingSpec{
		HavingType:  GreaterThan,
		Aggregation: "a",
		Value:       1000,
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
		"type": "greaterThan",
		"aggregation": "a",
		"value": 1000 
    }`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestDimSelectorHavingSpecSerialization(t *testing.T) {
	having := &DimSelectorHavingSpec{
		HavingType: DimSelector,
		Dimension:  "a",
		Value:      1000,
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
		"type": "dimSelector",
		"dimension": "a",
		"value": 1000 
    }`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestAndHavingSpecSerialization(t *testing.T) {
	having := &AndHavingSpec{
		HavingType: And,
		HavingSpecs: []Having{
			&NumericHavingSpec{
				HavingType:  GreaterThan,
				Aggregation: "a",
				Value:       100,
			},
			&NumericHavingSpec{
				HavingType:  LessThan,
				Aggregation: "a",
				Value:       1000,
			},
		},
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
            "type": "and",
            "havingSpecs": [
                {
                    "type": "greaterThan",
                    "aggregation": "a",
                    "value": 100
                },
                {
                    "type": "lessThan",
                    "aggregation": "a",
                    "value": 1000 
                }
            ]
        }`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestOrHavingSpecSerialization(t *testing.T) {
	having := &OrHavingSpec{
		HavingType: Or,
		HavingSpecs: []Having{
			&NumericHavingSpec{
				HavingType:  GreaterThan,
				Aggregation: "a",
				Value:       100,
			},
			&NumericHavingSpec{
				HavingType:  LessThan,
				Aggregation: "a",
				Value:       1000,
			},
		},
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
            "type": "or",
            "havingSpecs": [
                {
                    "type": "greaterThan",
                    "aggregation": "a",
                    "value": 100
                },
                {
                    "type": "lessThan",
                    "aggregation": "a",
                    "value": 1000 
                }
            ]
        }`

	require.JSONEq(t, expected, string(jsonBytes))
}

func TestNotHavingSpecSerialization(t *testing.T) {
	having := &NotHavingSpec{
		HavingType: Not,
		HavingSpec: &NumericHavingSpec{
			HavingType:  EqualTo,
			Aggregation: "a",
			Value:       1000,
		},
	}
	jsonBytes, _ := json.Marshal(having)

	expected := `{
        "type": "not",
        "havingSpec":
            {
                "type": "equalTo",
                "aggregation": "a",
                "value": 1000
            }
        }`

	require.JSONEq(t, expected, string(jsonBytes))
}
