package datasource

import "bytes"

// DataSourceType for druid query
type DataSourceType int

const (
	TABLE DataSourceType = iota
	QUERY
	UNION
	LOOKUP
)

var datasourceTypeStrings = []string{
	"table", "query", "union", "lookup",
}

func (d DataSourceType) name() string {
	return datasourceTypeStrings[d]
}

func (d DataSourceType) ordinal() int {
	return int(d)
}

func (d DataSourceType) values() *[]string {
	return &datasourceTypeStrings
}

func (d DataSourceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(d.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
