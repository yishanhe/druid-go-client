package datasource

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

// DataSourceType for druid query
type DataSourceType int

const (
	Table DataSourceType = iota
	Query
	Union
	Lookup
)

var datasourceTypeStrings = []string{
	"table", "query", "union", "lookup",
}

func (d DataSourceType) Name() string {
	return datasourceTypeStrings[d]
}

func (d DataSourceType) Ordinal() int {
	return int(d)
}

func (d DataSourceType) Values() *[]string {
	return &datasourceTypeStrings
}

func (d DataSourceType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(d.Name())
}
