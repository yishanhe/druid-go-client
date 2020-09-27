package datasource

type DataSource interface {
	Type() string
}

type SimpleDataSource string

func (s SimpleDataSource) Type() string {
	return Table.Name()
}

type QueryDataSource struct {
	DataSourceType DataSourceType `json:"type"`
}

func (m *QueryDataSource) Type() string {
	return Query.Name()
}

type TableDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	Name           string         `json:"name"`
}

func (m *TableDataSource) Type() string {
	return Table.Name()
}

// UnionDataSource unions table data sources
type UnionDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	DataSources    []string       `json:"dataSources"`
}

func (m *UnionDataSource) Type() string {
	return Union.Name()
}

type LookupDataSource struct {
	DataSourceType DataSourceType `json:"type"`
	Lookup         string         `json:"lookup"`
}

func (m *LookupDataSource) Type() string {
	return Lookup.Name()
}
