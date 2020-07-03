package extraction

import (
	"encoding/json"
	"testing"
)

func TestMarshalExtraction(t *testing.T) {
	var extraFn ExtractionFn
	extraFn = &RegexExtractionFn{ExtractionType: PARTIAL}


	jsonBytes, err := json.Marshal(extraFn)
	t.Log(string(jsonBytes))
	t.Log(err)
}

