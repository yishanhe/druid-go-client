package having

import "github.com/yishanhe/druid-go-client/query/components/filter"

type Having interface {
}

type FilterHavingSpec struct {
	HavingType HavingType    `json:"type"`
	Filter     filter.Filter `json:"filter"`
}

// equalTo, lessThan, greaterThan
type NumericHavingSpec struct {
	HavingType  HavingType `json:"type"`
	Aggregation string     `json:"aggregation"`
	Value       int64      `json:"value"`
}

type DimSelectorHavingSpec struct {
	HavingType HavingType `json:"type"`
	Dimension  string     `json:"dimension"`
	Value      int64      `json:"value"`
}

type AndHavingSpec struct {
	HavingType  HavingType `json:"type"`
	HavingSpecs []Having   `json:"havingSpecs"`
}

type OrHavingSpec struct {
	HavingType  HavingType `json:"type"`
	HavingSpecs []Having   `json:"havingSpecs"`
}

type NotHavingSpec struct {
	HavingType HavingType `json:"type"`
	HavingSpec Having     `json:"havingSpec"`
}
