package main

import (
	"fmt"

	"github.com/yishanhe/druid-go-client"
	"github.com/yishanhe/druid-go-client/query"
	"github.com/yishanhe/druid-go-client/query/components/datasource"
)

func main() {
	client := &druid.Client{URL: "http://127.0.0.1:8082/druid/v2/?pretty"}
	scanQuery := &query.ScanQuery{
		QueryType:  query.Scan,
		DataSource: datasource.SimpleDataSource("a"),
		Columns:    []string{"a", "a"},
		Intervals:  []string{"2019-01-01/2021-01-01"},
		Limit:      1,
	}
	resp, err := client.Query(scanQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

}
