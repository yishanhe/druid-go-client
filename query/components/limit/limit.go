package limit

import "github.com/yishanhe/druid-go-client/query/components/sorting"

type Limit interface {
	Type() string
}

type SimpleLimit struct {
	LimitType LimitType `json:"type"`
	Limit     int64     `json:"limit,omitempty"`
	Offset    int64     `json:"offset,omitempty"`
	Columns   []string  `json:"columns,omitempty"`
}

func (s SimpleLimit) Type() string {
	return s.LimitType.Name()
}

type DefaultLimit struct {
	LimitType LimitType               `json:"type"`
	Limit     int64                   `json:"limit,omitempty"`
	Offset    int64                   `json:"offset,omitempty"`
	Columns   []sorting.OrderByColumn `json:"columns,omitempty"`
}

func (d DefaultLimit) Type() string {
	return d.LimitType.Name()
}
