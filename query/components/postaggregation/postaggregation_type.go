package postaggregation

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type PostAggregationType int

const (
	Arithmetic PostAggregationType = iota
	FieldAccess
	FinalizingFieldAccess
	Constant
	DoubleGreatest
	LongGreatest
	DoubleLeast
	LongLeast
	Javascript
	HyperUniqueCardinality
)

var postAggregationTypeStrings = []string{
	"arithmetic", "fieldAccess", "finalizingFieldAccess", "constant", "doubleGreatest", "longGreatest", "doubleLeast", "longLeast",
	"javascript", "hyperUniqueCardinality",
}

func (p PostAggregationType) Name() string {
	return postAggregationTypeStrings[p]
}

func (p PostAggregationType) Ordinal() int {
	return int(p)
}

func (p PostAggregationType) Values() *[]string {
	return &postAggregationTypeStrings
}

func (p PostAggregationType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(p.Name())
}
