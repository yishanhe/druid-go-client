package virtualcolumn

import "github.com/yishanhe/druid-go-client/query/components/common"

type VirtualColumn interface {
	Type() string
}

type ExpressionVirtualColumn struct {
	VirtualColumnType VirtualColumnType `json:"type"`
	Name              string            `json:"name"`
	Expression        string            `json:"expression"`
	OutputType        common.OutputType `json:"outputType"`
}

func (e ExpressionVirtualColumn) Type() string {
	return e.VirtualColumnType.Name()
}
