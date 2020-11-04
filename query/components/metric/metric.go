package metric

import "github.com/yishanhe/druid-go-client/query/components/sorting"

type Metric interface {
	Type() string
}

type SimpleTopNMetric string

func (s SimpleTopNMetric) Type() string {
	return "simple"
}

type NumericTopNMetric struct {
	MetricType MetricType `json:"type"`
	Metric     Metric     `json:"metric"`
}

func (n NumericTopNMetric) Type() string {
	return NUMERIC.Name()
}

type DimensionTopNMetric struct {
	MetricType   MetricType           `json:"type"`
	Ordering     sorting.SortingOrder `json:"ordering,omitempty"`
	PreviousStop string               `json:"previousStop,omitempty"`
}

func (d DimensionTopNMetric) Type() string {
	return DIMENSION.Name()
}

type InvertedTopNMetric struct {
	MetricType MetricType `json:"type"`
	Metric     Metric     `json:"metric"`
}
