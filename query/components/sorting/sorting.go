package sorting

type Orderby interface {
	Type() string
}

type OrderByColumn struct {
	Dimension      string           `json:"dimension"`
	Direction      Direction        `json:"direction"`
	DimensionOrder SortingOrderType `json:"dimensionOrder"`
}

func (o OrderByColumn) Type() string {
	return o.DimensionOrder.Name()
}
