package query

import (
	"github.com/yishanhe/druid-go-client/query/components/aggregation"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
	"github.com/yishanhe/druid-go-client/query/components/postaggregation"
)

type Query interface {
	Type() string
}

type TimeseriesQuery struct {
	QueryType        QueryType                        `json:"queryType"`
	DataSource       datasource.DataSource            `json:"dataSource"`
	Granularity      granularity.Granularity          `json:"granularity"`
	Intervals        []string                         `json:"intervals"`
	Filter           filter.Filter                    `json:"filter,omitempty"`
	Aggregations     []aggregation.Aggregator         `json:"aggregations,omitempty"`
	PostAggregations []postaggregation.PostAggregator `json:"postAggregations,omitempty"`
	Context          map[string]interface{}           `json:"context,omitempty"`
	Limit            int64                            `json:"limit,omitempty"`
	Descending       bool                             `json:"descending,omitempty"`
}

type GroupByQuery struct {
	QueryType        QueryType                        `json:"queryType"`
	DataSource       datasource.DataSource            `json:"dataSource"`
	Granularity      granularity.Granularity          `json:"granularity"`
	Dimensions       []dimension.Dimension            `json:"dimensions"`
	Intervals        []string                         `json:"intervals"`
	Filter           filter.Filter                    `json:"filter,omitempty"`
	Aggregations     []aggregation.Aggregator         `json:"aggregations,omitempty"`
	PostAggregations []postaggregation.PostAggregator `json:"postAggregations,omitempty"`
	Context          map[string]interface{}           `json:"context,omitempty"`
}

func (g GroupByQuery) Type() string {
	return GroupBy.Name()
}

type GroupByResult struct {
	Version   string                 `json:"version"`
	Timestamp string                 `json:"timestamp"`
	Event     map[string]interface{} `json:"event"`
}

type ScanQuery struct {
	QueryType    QueryType             `json:"queryType"`
	DataSource   datasource.DataSource `json:"dataSource"`
	ResultFormat ResultFormat          `json:"resultFormat,omitempty"`
	Columns      []string              `json:"columns"`
	Intervals    []string              `json:"intervals"`
	BatchSize    int64                 `json:"batchSize,omitempty"`
	Limit        int64                 `json:"limit,omitempty"`
}

func (s ScanQuery) Type() string {
	return Scan.Name()
}

type ScanQueryResult struct {
	SegmentID string                   `json:"segmentId"`
	Columns   []string                 `json:"columns"`
	Events    []map[string]interface{} `json:"events"`
}

type ScanQueryCompactedResult struct {
	SegmentID string     `json:"segmentId"`
	Columns   []string   `json:"columns"`
	Events    [][]string `json:"events"`
}
