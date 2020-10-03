package virtualcolumn

import "github.com/yishanhe/druid-go-client/pkg/enum"

type VirtualColumnType int

const (
	Expression VirtualColumnType = iota
)

var virtualColumnTypeStrings = []string{"expression"}

func (v VirtualColumnType) Name() string {
	return virtualColumnTypeStrings[v]
}

func (v VirtualColumnType) Ordinal() int {
	return int(v)
}

func (v VirtualColumnType) Values() *[]string {
	return &virtualColumnTypeStrings
}

func (v VirtualColumnType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(v.Name())
}
