package query

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

// QueryType for druid query
type QueryType int

const (
	TimeSeries QueryType = iota
	TopN
	GroupBy
	Scan
	search
	TimeBoundary
	SegmentMetadata
	DataSoureMetadata
)

var queryTypeStrings = []string{
	"timeseries", "topN", "groupBy", "scan", "search", "timeBoundary", "segmentMetadata", "dataSourceMetadata",
}

func (q QueryType) Name() string {
	return queryTypeStrings[q]
}

func (q QueryType) Ordinal() int {
	return int(q)
}

func (q QueryType) Values() *[]string {
	return &queryTypeStrings
}

func (q QueryType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(q.Name())
}
