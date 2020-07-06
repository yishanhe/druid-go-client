package datasource

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleDataSource(t *testing.T) {
	var m DataSource
	m = SimpleDataSource("dataSource1")
	jsonBytes, _ := json.Marshal(m)
	result := string(jsonBytes)
	expected := `"dataSource1"`
	assert.Equal(t, expected, result)
}

func TestQueryDataSource(t *testing.T) {

	var m DataSource
	// var m models.QueryDataSource
	m = &QueryDataSource{
		DataSourceType: QUERY,
	}

	jsonBytes, _ := json.Marshal(m)
	result := string(jsonBytes)
	t.Log(result)
}
