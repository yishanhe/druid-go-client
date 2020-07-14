package aggregation

import "bytes"

type AggregationType int

const (
	COUNT AggregationType = iota
	LONG_SUM
	DOUBLE_SUM
	FLOAT_SUM
	DOUBLE_MIN
	DOUBLE_MAX
	FLOAT_MIN
	FLOAT_MAX
	LONG_MIN
	LONG_MAX
	DOUBLE_MEAN
	DOUBLE_FIRST
	DOUBLE_LAST
	FLOAT_FIRST
	FLOAT_LAST
	LONG_FIRST
	LONG_LAST
	STRING_FIRST
	STRING_LAST
	DOUBLE_ANY
	FLOAT_ANY
	LONG_ANY
	STRING_ANY
	JAVASCRIPT
	THETA_SKETCH
	FILTERED
)

var aggregationTypeStrings = []string{
	"count", "longSum", "doubleSum", "floatSum", "doubleMin", "doubleMax", "floatMin", "floatMax", "longMin", "longMax",
	"doubleMean", "doubleFirst", "doubleLast", "floatFirst", "floatLast", "longFirst", "longLast", "stringFirst", "stringLast",
	"doubleAny", "floatAny", "longAny", "stringAny", "javascript", "thetaSketch", "filtered",
}

func (a AggregationType) name() string {
	return aggregationTypeStrings[a]
}

func (a AggregationType) ordinal() int {
	return int(a)
}

func (a AggregationType) values() *[]string {
	return &aggregationTypeStrings
}

func (a AggregationType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(a.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

