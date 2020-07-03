package main

import (
	"fmt"

	query2 "github.com/yishanhe/druid-go-client/query"

	"github.com/yishanhe/druid-go-client/query/components/granularity"

	"github.com/yishanhe/druid-go-client"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
)

func main() {
	client := &druid.Client{URL: "http://127.0.0.1:8082/druid/v2/?pretty"}
	query := &query2.GroupByQuery{
		QueryType:  query2.GROUPBY,
		DataSource: datasource.SimpleDataSource("correlated_flow_viz"),
		Dimensions: []dimension.Dimension{
			dimension.SimpleDimension("srcVmId"),
		},
		Granularity: granularity.HOUR,
		Intervals:   []string{"2019-01-01/2021-01-01"},
	}
	resp, err := client.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

}
