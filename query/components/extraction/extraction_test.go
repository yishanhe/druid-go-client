package extraction

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/lookup"
)

func TestTimeFormatExtractionSerialization(t *testing.T) {
	extraFn := &TimeFormatExtraction{
		ExtractionType: TimeFormat,
		Format:         "EEEE",
		TimeZone:       "America/Montreal",
		Locale:         "fr",
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{
		"type" : "timeFormat",
		"format" : "EEEE",
		"timeZone" : "America/Montreal",
		"locale" : "fr"
	  }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestStrlenExtractionSerialization(t *testing.T) {
	extraFn := &StrlenExtraction{
		ExtractionType: Strlen,
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{ "type" : "strlen" }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestSubstringExtractionSerialization(t *testing.T) {
	extraFn := &SubstringExtraction{
		ExtractionType: Substring,
		Index:          1,
		Length:         4,
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{ "type" : "substring", "index" : 1, "length" : 4 }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestJavascriptExtractionSerialization(t *testing.T) {
	extraFn := &JavascriptExtraction{
		ExtractionType: Javascript,
		Function:       "function(str) { return str + '!!!'; }",
		Injective:      true,
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{
		"type" : "javascript",
		"function" : "function(str) { return str + '!!!'; }",
		"injective" : true
	  }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestRegisteredLookupExtractionSerialization(t *testing.T) {
	extraFn := &RegisteredLookupExtraction{
		ExtractionType:     RegisteredLookup,
		Lookup:             "some_lookup_name",
		RetainMissingValue: true,
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{
		"type": "registeredLookup",
		"lookup": "some_lookup_name",
		"retainMissingValue": true
	  }`
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestLookupExtractionSerialization(t *testing.T) {
	extraFn := &LookupExtraction{
		ExtractionType: Lookup,
		Lookup: &lookup.MapLookup{
			LookupType: lookup.Map,
			Map: map[string]string{
				"foo": "bar",
				"baz": "bat",
			},
		},
		RetainMissingValue:      true,
		Injective:               true,
		ReplaceMissingValueWith: "MISSING",
	}
	jsonBytes, _ := json.Marshal(extraFn)
	expected := `{
		"type":"lookup",
		"lookup":{
		  "type":"map",
		  "map":{"foo":"bar", "baz":"bat"}
		},
		"retainMissingValue":true,
		"injective":true,
		"replaceMissingValueWith":"MISSING"
	  }`
	require.JSONEq(t, expected, string(jsonBytes))
}
