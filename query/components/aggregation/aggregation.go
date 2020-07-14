package aggregation

import "github.com/yishanhe/druid-go-client/query/components/filter"

type Aggregator interface {
	Type() string
}

type Aggregation struct {
	AggregationType AggregationType `json:"type"`
	Name string `json:"name"`
	FieldName string `json:"fieldName"`
	MaxStringBytes int `json:"maxStringBytes,omitempty"`
	AggregationFn string `json:"fnAggregate"`
	CombineFn string `json:"fnCombine"`
	ResetFn string `json:"fnReset"`
}

func (s Aggregation) Type() string {
	return s.AggregationType.name()
}

type FilteredAggregation struct {
	AggregationType AggregationType `json:"type"`	
	Aggregator Aggregator `json:"aggregator"`
	Filter  filter.Filter `json:"filter"`
}

func (f FilteredAggregation) Type() string {
	return FILTERED.name()
}





