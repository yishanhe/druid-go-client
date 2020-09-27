package aggregation

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type AggregationType int

const (
	Count AggregationType = iota
	LongSum
	DoubleSum
	FloatSum
	DoubleMin
	DoubleMax
	FloatMin
	FloatMax
	LongMin
	LongMax
	DoubleMean
	DoubleFirst
	DoubleLast
	FloatFirst
	FloatLast
	LongFirst
	LongLast
	StringFirst
	StringLast
	DoubleAny
	FloatAny
	LongAny
	StringAny
	Javascript
	ThetaSketch
	Filtered
)

var aggregationTypeStrings = []string{
	"count", "longSum", "doubleSum", "floatSum", "doubleMin", "doubleMax", "floatMin", "floatMax", "longMin", "longMax",
	"doubleMean", "doubleFirst", "doubleLast", "floatFirst", "floatLast", "longFirst", "longLast", "stringFirst", "stringLast",
	"doubleAny", "floatAny", "longAny", "stringAny", "javascript", "thetaSketch", "filtered",
}

func (a AggregationType) Name() string {
	return aggregationTypeStrings[a]
}

func (a AggregationType) Ordinal() int {
	return int(a)
}

func (a AggregationType) Values() *[]string {
	return &aggregationTypeStrings
}

func (a AggregationType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(a.Name())
}
