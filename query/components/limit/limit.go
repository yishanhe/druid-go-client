package limit

import "github.com/yishanhe/druid-go-client/query/components/sorting"

type Limit interface {
	Type() string
}

type DefaultLimit struct {
	LimitType LimitType               `json:"type"`
	Limit     int64                   `json:"limit"`
	Columns   []sorting.OrderByColumn `json:"columns"`
}

func (d DefaultLimit) Type() string {
	return d.LimitType.Name()
}
