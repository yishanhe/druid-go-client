package postaggregation

import "bytes"

type PostAggregationType int

const (
	ARITHMETIC PostAggregationType = iota
	FIELD_ACCESS
	FINALIZING_FIELD_ACCESS
	CONSTANT
	DOUBLE_GREATEST
	LONG_GREATEST
	DOUBLE_LEAST
	LONG_LEAST
	JAVASCRIPT
	HYPER_UNIQUE_CARDINALITY
)

var postAggregationTypeStrings = []string{
	"arithmetic", "fieldAcess", "finalizingFieldAccess", "constant", "doubleGreatest", "longGreatest", "doubleLeast", "longLeast",
	"javascript", "hyperUniqueCardinality",
}

func (p PostAggregationType) name() string {
	return postAggregationTypeStrings[p]
}

func (p PostAggregationType) ordinal() int {
	return int(p)
}

func (p PostAggregationType) values() *[]string {
	return &postAggregationTypeStrings
}

func (p PostAggregationType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(p.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}



