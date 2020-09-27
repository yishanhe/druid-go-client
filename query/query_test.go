package query

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
)

func TestGroupByQuery(t *testing.T) {
	query := &GroupByQuery{
		QueryType:  GroupBy,
		DataSource: datasource.SimpleDataSource("testDataSource"),
		Dimensions: []dimension.Dimension{
			dimension.SimpleDimension("abc"),
		},
		Intervals:   []string{"2020-07-06T12:21:21/2020-07-06T12:21:21"},
		Granularity: granularity.DAY,
	}
	expected := `{
		"queryType": "groupBy",
        "intervals": ["2020-07-06T12:21:21/2020-07-06T12:21:21"],
		"dataSource": "testDataSource",
        "dimensions": ["abc"],
		"granularity": "day"
    }`
	jsonBytes, _ := json.Marshal(query)
	require.JSONEq(t, expected, string(jsonBytes))
}
