package query

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/aggregation"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/filter"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
	"github.com/yishanhe/druid-go-client/query/components/having"
	"github.com/yishanhe/druid-go-client/query/components/limit"
	"github.com/yishanhe/druid-go-client/query/components/metric"
	"github.com/yishanhe/druid-go-client/query/components/postaggregation"
	"github.com/yishanhe/druid-go-client/query/components/search"
	"github.com/yishanhe/druid-go-client/query/components/sorting"
)

func TestTimeseriesQuerySerialization(t *testing.T) {
	query := TimeseriesQuery{
		QueryType:   TimeSeries,
		DataSource:  datasource.SimpleDataSource("sample_datasource"),
		Granularity: granularity.DAY,
		Descending:  true,
		Filter: filter.AndFilter{
			FilterType: filter.And,
			Fields: []filter.Filter{
				filter.SelectorFilter{
					FilterType: filter.Selector,
					Dimension:  "sample_dimension1",
					Value:      "sample_value1",
				},
				filter.OrFilter{
					FilterType: filter.Or,
					Fields: []filter.Filter{
						filter.SelectorFilter{
							FilterType: filter.Selector,
							Dimension:  "sample_dimension2",
							Value:      "sample_value2",
						},
						filter.SelectorFilter{
							FilterType: filter.Selector,
							Dimension:  "sample_dimension3",
							Value:      "sample_value3",
						},
					},
				},
			},
		},
		Aggregations: []aggregation.Aggregator{
			aggregation.Aggregation{
				AggregationType: aggregation.LongSum,
				Name:            "sample_name1",
				FieldName:       "sample_fieldName1",
			},
			aggregation.Aggregation{
				AggregationType: aggregation.DoubleSum,
				Name:            "sample_name2",
				FieldName:       "sample_fieldName2",
			},
		},
		PostAggregations: []postaggregation.PostAggregator{
			postaggregation.ArithmeticPostAggregation{
				PostAggregationType: postaggregation.Arithmetic,
				Name:                "sample_divide",
				Function:            "/",
				Fields: []postaggregation.PostAggregator{
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						Name:                "postAgg__sample_name1",
						FieldName:           "sample_name1",
					},
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						Name:                "postAgg__sample_name2",
						FieldName:           "sample_name2",
					},
				},
			},
		},
		Intervals: []string{"2012-01-01T00:00:00.000/2012-01-03T00:00:00.000"},
	}
	expected := `{
		"queryType": "timeseries",
		"dataSource": "sample_datasource",
		"granularity": "day",
		"descending": true,
		"filter": {
		  "type": "and",
		  "fields": [
			{ "type": "selector", "dimension": "sample_dimension1", "value": "sample_value1" },
			{ "type": "or",
			  "fields": [
				{ "type": "selector", "dimension": "sample_dimension2", "value": "sample_value2" },
				{ "type": "selector", "dimension": "sample_dimension3", "value": "sample_value3" }
			  ]
			}
		  ]
		},
		"aggregations": [
		  { "type": "longSum", "name": "sample_name1", "fieldName": "sample_fieldName1" },
		  { "type": "doubleSum", "name": "sample_name2", "fieldName": "sample_fieldName2" }
		],
		"postAggregations": [
		  { "type": "arithmetic",
			"name": "sample_divide",
			"fn": "/",
			"fields": [
			  { "type": "fieldAccess", "name": "postAgg__sample_name1", "fieldName": "sample_name1" },
			  { "type": "fieldAccess", "name": "postAgg__sample_name2", "fieldName": "sample_name2" }
			]
		  }
		],
		"intervals": [ "2012-01-01T00:00:00.000/2012-01-03T00:00:00.000" ]
	  }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestTopNQuerySeralization(t *testing.T) {
	query := TopNQuery{
		QueryType:   TopN,
		DataSource:  datasource.SimpleDataSource("sample_data"),
		Dimension:   dimension.SimpleDimension("sample_dim"),
		Threshold:   5,
		Metric:      metric.SimpleTopNMetric("count"),
		Granularity: granularity.ALL,
		Filter: filter.AndFilter{
			FilterType: filter.And,
			Fields: []filter.Filter{
				filter.SelectorFilter{
					FilterType: filter.Selector,
					Dimension:  "dim1",
					Value:      "some_value",
				},
				filter.SelectorFilter{
					FilterType: filter.Selector,
					Dimension:  "dim2",
					Value:      "some_other_val",
				},
			},
		},
		Aggregations: []aggregation.Aggregator{
			aggregation.Aggregation{
				AggregationType: aggregation.LongSum,
				Name:            "count",
				FieldName:       "count",
			},
			aggregation.Aggregation{
				AggregationType: aggregation.DoubleSum,
				Name:            "some_metric",
				FieldName:       "some_metric",
			},
		},
		PostAggregations: []postaggregation.PostAggregator{
			postaggregation.ArithmeticPostAggregation{
				PostAggregationType: postaggregation.Arithmetic,
				Name:                "average",
				Function:            "/",
				Fields: []postaggregation.PostAggregator{
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						Name:                "some_metric",
						FieldName:           "some_metric",
					},
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						Name:                "count",
						FieldName:           "count",
					},
				},
			},
		},
		Intervals: []string{"2013-08-31T00:00:00.000/2013-09-03T00:00:00.000"},
	}
	expected := `{
		"queryType": "topN",
		"dataSource": "sample_data",
		"dimension": "sample_dim",
		"threshold": 5,
		"metric": "count",
		"granularity": "all",
		"filter": {
		  "type": "and",
		  "fields": [
			{
			  "type": "selector",
			  "dimension": "dim1",
			  "value": "some_value"
			},
			{
			  "type": "selector",
			  "dimension": "dim2",
			  "value": "some_other_val"
			}
		  ]
		},
		"aggregations": [
		  {
			"type": "longSum",
			"name": "count",
			"fieldName": "count"
		  },
		  {
			"type": "doubleSum",
			"name": "some_metric",
			"fieldName": "some_metric"
		  }
		],
		"postAggregations": [
		  {
			"type": "arithmetic",
			"name": "average",
			"fn": "/",
			"fields": [
			  {
				"type": "fieldAccess",
				"name": "some_metric",
				"fieldName": "some_metric"
			  },
			  {
				"type": "fieldAccess",
				"name": "count",
				"fieldName": "count"
			  }
			]
		  }
		],
		"intervals": [
		  "2013-08-31T00:00:00.000/2013-09-03T00:00:00.000"
		]
	  }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestGroupByQuerySerialization(t *testing.T) {
	query := &GroupByQuery{
		QueryType:   GroupBy,
		DataSource:  datasource.SimpleDataSource("sample_datasource"),
		Granularity: granularity.DAY,
		Dimensions: []dimension.Dimension{
			dimension.SimpleDimension("country"),
			dimension.SimpleDimension("device"),
		},
		LimitSpec: limit.SimpleLimit{
			LimitType: limit.Default,
			Limit:     5000,
			Columns:   []string{"country", "data_transfer"},
		},
		Filter: filter.AndFilter{
			FilterType: filter.And,
			Fields: []filter.Filter{
				filter.SelectorFilter{
					FilterType: filter.Selector,
					Dimension:  "carrier",
					Value:      "AT&T",
				},
				filter.OrFilter{
					FilterType: filter.Or,
					Fields: []filter.Filter{
						filter.SelectorFilter{
							FilterType: filter.Selector,
							Dimension:  "make",
							Value:      "Apple",
						},
						filter.SelectorFilter{
							FilterType: filter.Selector,
							Dimension:  "make",
							Value:      "Samsung",
						},
					},
				},
			},
		},
		Aggregations: []aggregation.Aggregator{
			aggregation.Aggregation{
				AggregationType: aggregation.LongSum,
				Name:            "total_usage",
				FieldName:       "user_count",
			},
			aggregation.Aggregation{
				AggregationType: aggregation.DoubleSum,
				Name:            "data_transfer",
				FieldName:       "data_transfer",
			},
		},
		PostAggregations: []postaggregation.PostAggregator{
			postaggregation.ArithmeticPostAggregation{
				PostAggregationType: postaggregation.Arithmetic,
				Name:                "avg_usage",
				Function:            "/",
				Fields: []postaggregation.PostAggregator{
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						FieldName:           "data_transfer",
					},
					postaggregation.FieldAccessPostAggregation{
						PostAggregationType: postaggregation.FieldAccess,
						FieldName:           "total_usage",
					},
				},
			},
		},
		Intervals: []string{"2012-01-01T00:00:00.000/2012-01-03T00:00:00.000"},
		Having: having.NumericHavingSpec{
			HavingType:  having.GreaterThan,
			Aggregation: "total_usage",
			Value:       100,
		},
	}
	expected := `{
		"queryType": "groupBy",
		"dataSource": "sample_datasource",
		"granularity": "day",
		"dimensions": ["country", "device"],
		"limitSpec": { "type": "default", "limit": 5000, "columns": ["country", "data_transfer"] },
		"filter": {
		  "type": "and",
		  "fields": [
			{ "type": "selector", "dimension": "carrier", "value": "AT&T" },
			{ "type": "or",
			  "fields": [
				{ "type": "selector", "dimension": "make", "value": "Apple" },
				{ "type": "selector", "dimension": "make", "value": "Samsung" }
			  ]
			}
		  ]
		},
		"aggregations": [
		  { "type": "longSum", "name": "total_usage", "fieldName": "user_count" },
		  { "type": "doubleSum", "name": "data_transfer", "fieldName": "data_transfer" }
		],
		"postAggregations": [
		  { "type": "arithmetic",
			"name": "avg_usage",
			"fn": "/",
			"fields": [
			  { "type": "fieldAccess", "fieldName": "data_transfer" },
			  { "type": "fieldAccess", "fieldName": "total_usage" }
			]
		  }
		],
		"intervals": [ "2012-01-01T00:00:00.000/2012-01-03T00:00:00.000" ],
		"having": {
		  "type": "greaterThan",
		  "aggregation": "total_usage",
		  "value": 100
		}
	  }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestScanQuerySerialization(t *testing.T) {
	query := ScanQuery{
		QueryType:    Scan,
		DataSource:   datasource.SimpleDataSource("wikipedia"),
		ResultFormat: List,
		Columns:      []string{},
		Intervals:    []string{"2013-01-01/2013-01-02"},
		BatchSize:    20480,
		Limit:        3,
	}
	expected := ` {
		"queryType": "scan",
		"dataSource": "wikipedia",
		"resultFormat": "list",
		"columns":[],
		"intervals": [
		  "2013-01-01/2013-01-02"
		],
		"batchSize":20480,
		"limit":3
	  }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}

func TestSearchQuerySerialization(t *testing.T) {
	query := SearchQuery{
		QueryType:        Search,
		DataSource:       datasource.SimpleDataSource("sample_datasource"),
		Granularity:      granularity.DAY,
		SearchDimensions: []string{"dim1", "dim2"},
		Query: search.InsensitiveContainsSearch{
			SearchType: search.InsensitiveContains,
			Value:      "Ke",
		},
		Sort: sorting.SortingOrder{
			SortingOrderType: sorting.Lexicographic,
		},
		Intervals: []string{"2013-01-01T00:00:00.000/2013-01-03T00:00:00.000"},
	}
	expected := `{
		"queryType": "search",
		"dataSource": "sample_datasource",
		"granularity": "day",
		"searchDimensions": [
		  "dim1",
		  "dim2"
		],
		"query": {
		  "type": "insensitive_contains",
		  "value": "Ke"
		},
		"sort" : {
		  "type": "lexicographic"
		},
		"intervals": [
		  "2013-01-01T00:00:00.000/2013-01-03T00:00:00.000"
		]
	  }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}
