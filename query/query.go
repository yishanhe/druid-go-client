package query

import (
	"github.com/yishanhe/druid-go-client/query/components/aggregation"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
	"github.com/yishanhe/druid-go-client/query/components/having"
	"github.com/yishanhe/druid-go-client/query/components/limit"
	"github.com/yishanhe/druid-go-client/query/components/metric"
	"github.com/yishanhe/druid-go-client/query/components/postaggregation"
)

type Query interface {
	Type() string
}

type TimeseriesQuery struct {
	QueryType        QueryType                        `json:"queryType"`
	DataSource       datasource.DataSource            `json:"dataSource"`
	Descending       bool                             `json:"descending,omitempty"`
	Intervals        []string                         `json:"intervals"`
	Granularity      granularity.Granularity          `json:"granularity"`
	Filter           filter.Filter                    `json:"filter,omitempty"`
	Aggregations     []aggregation.Aggregator         `json:"aggregations,omitempty"`
	PostAggregations []postaggregation.PostAggregator `json:"postAggregations,omitempty"`
	Limit            int64                            `json:"limit,omitempty"`
	Context          map[string]interface{}           `json:"context,omitempty"`
}

func (t TimeseriesQuery) Type() string {
	return TimeSeries.Name()
}

type TimeseriesQueryResult struct {
	Timestamp string                 `json:"timestamp"`
	Result    map[string]interface{} `json:"result"`
}

type TopNQuery struct {
	QueryType        QueryType                        `json:"queryType"`
	DataSource       datasource.DataSource            `json:"dataSource"`
	Intervals        []string                         `json:"intervals"`
	Granularity      granularity.Granularity          `json:"granularity"`
	Filter           filter.Filter                    `json:"filter,omitempty"`
	Aggregations     []aggregation.Aggregator         `json:"aggregations,omitempty"`
	PostAggregations []postaggregation.PostAggregator `json:"postAggregations,omitempty"`
	Dimensions       []dimension.Dimension            `json:"dimensions"`
	Threshold        int64                            `json:"threshold"`
	Metric           metric.Metric                    `json:"metric"`
	Context          map[string]interface{}           `json:"context,omitempty"`
}

func (t TopNQuery) Type() string {
	return TopN.Name()
}

type TopNQueryResult TimeseriesQueryResult

type GroupByQuery struct {
	QueryType        QueryType                        `json:"queryType"`
	DataSource       datasource.DataSource            `json:"dataSource"`
	Dimensions       []dimension.Dimension            `json:"dimensions"`
	LimitSpec        limit.Limit                      `json:"limitSpec,omitempty"`
	Having           having.Having                    `json:"having,omitempty"`
	Granularity      granularity.Granularity          `json:"granularity"`
	Filter           filter.Filter                    `json:"filter,omitempty"`
	Aggregations     []aggregation.Aggregator         `json:"aggregations,omitempty"`
	PostAggregations []postaggregation.PostAggregator `json:"postAggregations,omitempty"`
	Intervals        []string                         `json:"intervals"`
	SubtotalsSpec    [][]string                       `json:"subtotalsSpec,omitempty"`
	Context          map[string]interface{}           `json:"context,omitempty"`
}

func (g GroupByQuery) Type() string {
	return GroupBy.Name()
}

type GroupByQueryResult struct {
	Version   string                 `json:"version"`
	Timestamp string                 `json:"timestamp"`
	Event     map[string]interface{} `json:"event"`
}

type ScanQuery struct {
	QueryType    QueryType             `json:"queryType"`
	DataSource   datasource.DataSource `json:"dataSource"`
	Intervals    []string              `json:"intervals"`
	ResultFormat ResultFormat          `json:"resultFormat,omitempty"`
	Columns      []string              `json:"columns"`
	BatchSize    int64                 `json:"batchSize,omitempty"`
	Limit        int64                 `json:"limit,omitempty"`
	Filter       filter.Filter         `json:"filter,omitempty"`
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
