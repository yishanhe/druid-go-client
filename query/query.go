package query

import (
	"github.com/yishanhe/druid-go-client/query/components/aggregation"
	"github.com/yishanhe/druid-go-client/query/components/common"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
	"github.com/yishanhe/druid-go-client/query/components/having"
	"github.com/yishanhe/druid-go-client/query/components/limit"
	"github.com/yishanhe/druid-go-client/query/components/metric"
	"github.com/yishanhe/druid-go-client/query/components/postaggregation"
	"github.com/yishanhe/druid-go-client/query/components/search"
	"github.com/yishanhe/druid-go-client/query/components/segment"
	"github.com/yishanhe/druid-go-client/query/components/sorting"
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
	Dimension        dimension.Dimension              `json:"dimension"`
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
	QueryType    QueryType              `json:"queryType"`
	DataSource   datasource.DataSource  `json:"dataSource"`
	Intervals    []string               `json:"intervals"`
	ResultFormat ResultFormat           `json:"resultFormat"`
	Filter       filter.Filter          `json:"filter,omitempty"`
	Columns      []string               `json:"columns"`
	BatchSize    int64                  `json:"batchSize,omitempty"`
	Limit        int64                  `json:"limit,omitempty"`
	Offset       int64                  `json:"offset,omitempty"`
	Order        sorting.Direction      `json:"order,omitempty"`
	Legacy       bool                   `json:"legacy,omitempty"`
	Context      map[string]interface{} `json:"context,omitempty"`
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

type SearchQuery struct {
	QueryType        QueryType               `json:"queryType"`
	DataSource       datasource.DataSource   `json:"dataSource"`
	Granularity      granularity.Granularity `json:"granularity"`
	Filter           filter.Filter           `json:"filter,omitempty"`
	Limit            int64                   `json:"limit,omitempty"`
	Intervals        []string                `json:"intervals"`
	SearchDimensions []string                `json:"searchDimensions,omitempty"`
	Query            search.Search           `json:"query"`
	Sort             sorting.SortingOrder    `json:"sort,omitempty"`
	Context          map[string]interface{}  `json:"context,omitempty"`
}

type SearchQueryResult TimeseriesQueryResult

type TimeBoundaryQuery struct {
	QueryType  QueryType              `json:"queryType"`
	DataSource datasource.DataSource  `json:"dataSource"`
	Bound      common.Bound           `json:"bound,omitempty"`
	Filter     filter.Filter          `json:"filter,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

type TimeBoundaryQueryResult TimeseriesQueryResult

type SegmentMetaDataQuery struct {
	QueryType              QueryType              `json:"queryType"`
	DataSource             datasource.DataSource  `json:"dataSource"`
	Intervals              []string               `json:"intervals,omitempty"`
	ToInclude              segment.Include        `json:"toInclude,omitempty"`
	Merge                  bool                   `json:"merge"`
	Context                map[string]interface{} `json:"context,omitempty"`
	AnalysisTypes          []segment.AnalysisType `json:"analysisTypes,omitempty"`
	LenientAggregatorMerge bool                   `json:"lenientAggregatorMerge"`
}

type SegmentMetaDataQueryResult struct {
	Timestamp        string                 `json:"timestamp"`
	Intervals        []string               `json:"intervals"`
	Columns          map[string]interface{} `json:"columns"` // todo add type for column
	Aggregators      map[string]interface{} `json:"aggregators,omitempty"`
	Cardinality      interface{}            `json:"cardinality,omitempty"`
	Minmax           interface{}            `json:"minmax,omitempty"`
	Size             int64                  `json:"size,omitempty"`
	TimestampSpec    interface{}            `json:"timestampSpec,omitempty"`
	QueryGranularity interface{}            `json:"queryGranularity,omitempty"`
	Rollup           bool                   `json:"rollup,omitempty"`
}

type DatasourceMetadataQuery struct {
	QueryType  QueryType              `json:"queryType"`
	DataSource datasource.DataSource  `json:"dataSource"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

type DatasourceMetadataQueryResult TimeseriesQueryResult
