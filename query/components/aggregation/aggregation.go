package aggregation

import "github.com/yishanhe/druid-go-client/query/components/filter"

type Aggregator interface {
	Type() string
}

// all the rests aggregators are modeled here
type Aggregation struct {
	AggregationType AggregationType `json:"type"`
	Name            string          `json:"name"`
	FieldName       string          `json:"fieldName"`
	MaxStringBytes  int             `json:"maxStringBytes,omitempty"`
}

func (s Aggregation) Type() string {
	return s.AggregationType.Name()
}

type FilteredAggregation struct {
	AggregationType AggregationType `json:"type"`
	Aggregator      Aggregator      `json:"aggregator"`
	Filter          filter.Filter   `json:"filter"`
}

func (f FilteredAggregation) Type() string {
	return Filtered.Name()
}

type JavascriptAggregation struct {
	AggregationType AggregationType `json:"type"`
	Name            string          `json:"name"`
	FieldNames      []string        `json:"fieldNames"`
	AggregationFn   string          `json:"fnAggregate"`
	CombineFn       string          `json:"fnCombine"`
	ResetFn         string          `json:"fnReset"`
}

func (j JavascriptAggregation) Type() string {
	return Javascript.Name()
}
