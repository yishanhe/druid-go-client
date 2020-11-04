package metric

import "github.com/yishanhe/druid-go-client/pkg/enum"

type MetricType int

const (
	NUMERIC MetricType = iota
	DIMENSION
	INVERTED
)

var metricTypeStrings = []string{"numeric", "dimension", "inverted"}

func (t MetricType) Name() string {
	return metricTypeStrings[t]
}

func (t MetricType) Ordinal() int {
	return int(t)
}

func (t MetricType) Values() *[]string {
	return &metricTypeStrings
}

func (t MetricType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(t.Name())
}
