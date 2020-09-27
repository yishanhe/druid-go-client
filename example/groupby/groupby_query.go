package main

import (
	"fmt"

	"github.com/yishanhe/druid-go-client/query"

	"github.com/yishanhe/druid-go-client/query/components/granularity"

	"github.com/yishanhe/druid-go-client"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
)

func main() {
	client := &druid.Client{URL: "http://127.0.0.1:8082/druid/v2/?pretty"}
	groupByQuery := &query.GroupByQuery{
		QueryType:  query.GroupBy,
		DataSource: datasource.SimpleDataSource("example"),
		Dimensions: []dimension.Dimension{
			dimension.SimpleDimension("a"),
		},
		Granularity: granularity.HOUR,
		Intervals:   []string{"2019-01-01/2021-01-01"},
	}
	resp, err := client.Query(groupByQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

}
