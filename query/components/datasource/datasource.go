package datasource

type DataSource interface {
	Type() string
}

type SimpleDataSource string

func (s SimpleDataSource) Type() string {
	return TABLE.name()
}

type QueryDataSource struct {
	DataSourceType DataSourceType `json:"type"`
}

func (m *QueryDataSource) Type() string {
	return QUERY.name()
}

type TableDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	Name           string         `json:"name"`
}

func (m *TableDataSource) Type() string {
	return TABLE.name()
}

// UnionDataSource unions table data sources
type UnionDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	DataSources    []string       `json:"dataSources"`
}

func (m *UnionDataSource) Type() string {
	return UNION.name()
}

type LookupDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	Lookup         string         `json:"lookup"`
}

func (m *LookupDataSource) Type() string {
	return LOOKUP.name()
}
