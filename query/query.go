package query

import (
	"github.com/yishanhe/druid-go-client/query/components/datasource"
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/granularity"
)

type Query interface {
	Type() string
}

type GroupByQuery struct {
	QueryType   QueryType               `json:"queryType"`
	DataSource  datasource.DataSource   `json:"dataSource"`
	Dimensions  []dimension.Dimension   `json:"dimensions"`
	Intervals   []string                `json:"intervals"`
	Granularity granularity.Granularity `json:"granularity"`
}

func (g GroupByQuery) Type() string {
	return GROUPBY.name()
}

type GroupByResult struct {
	Version   string                 `json:"version"`
	Timestamp string                 `json:"timestamp"`
	Event     map[string]interface{} `json:"event"`
}
