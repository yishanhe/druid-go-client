package filter

import (
	"github.com/yishanhe/druid-go-client/query/components/dimension"
	"github.com/yishanhe/druid-go-client/query/components/extraction"
	"github.com/yishanhe/druid-go-client/query/components/sorting"
)

type Filter interface {
	Type() string
}

type ColumnComparisonFilter struct {
	FilterType FilterType            `json:"type"`
	Dimensions []dimension.Dimension `json:"dimensions"`
}

func (c ColumnComparisonFilter) Type() string {
	return COLUMN_COMPARISON.name()
}

type RegexFilter struct {
	FilterType FilterType `json:"type"`
	Dimension  string     `json:"dimension"`
	Pattern    string     `json:"pattern"`
}

func (r RegexFilter) Type() string {
	return REGEX.name()
}

type SelectorFilter struct {
	FilterType FilterType            `json:"type"`
	Dimension  string                `json:"dimension"`
	Value      string                `json:"value"`
	Extraction extraction.Extraction `json:"extractionFn,omitempty"`
}

func (s SelectorFilter) Type() string {
	return SELECTOR.name()
}

type AndFilter struct {
	FilterType FilterType `json:"type"`
	Fields     []Filter   `json:"fields"`
}

func (a AndFilter) Type() string {
	return AND.name()
}

type OrFilter struct {
	FilterType FilterType `json:"type"`
	Fields     []Filter   `json:"fields"`
}

func (o OrFilter) Type() string {
	return OR.name()
}

type NotFilter struct {
	FilterType FilterType `json:"type"`
	Field      Filter     `json:"field"`
}

func (n NotFilter) Type() string {
	return NOT.name()
}

type JavaScriptFilter struct {
	FilterType FilterType `json:"type"`
	Dimension  string     `json:"dimension"`
	Function   string     `json:"function"`
}

func (j JavaScriptFilter) Type() string {
	return JAVASCRIPT.name()
}

type ExtractionFilter struct {
	FilterType FilterType            `json:"type"`
	Dimension  string                `json:"dimension"`
	Value      string                `json:"value"`
	Extraction extraction.Extraction `json:"extractionFn"`
}

func (e ExtractionFilter) Type() string {
	return EXTRACTION.name()
}

type SearchFilter struct {
	FilterType FilterType            `json:"type"`
	Dimension  string                `json:"dimension"`
	Query      SearchQuerySpec       `json:"query"`
	Extraction extraction.Extraction `json:"extractionFn,omitempty"`
}

type SearchQuerySpec struct {
	Type          string `json:"type"`
	Value         string `json:"value"`
	CaseSensitive bool   `json:"caseSensitive,omitempty"`
}

func (s SearchFilter) Type() string {
	return SEARCH.name()
}

type InFilter struct {
	FilterType FilterType `json:"type"`
	Dimension  string     `json:"dimension"`
	Values     []string   `json:"values"`
}

func (i InFilter) Type() string {
	return IN.name()
}

type LikeFilter struct {
	FilterType FilterType            `json:"type"`
	Dimension  string                `json:"dimension"`
	Pattern    string                `json:"pattern"`
	Escape     string                `json:"escape,omitempty"`
	Extraction extraction.Extraction `json:"extractionFn,omitempty"`
}

func (l LikeFilter) Type() string {
	return LIKE.name()
}

type BoundFilter struct {
	FilterType  FilterType            `json:"type"`
	Dimension   string                `json:"dimension"`
	Lower       string                `json:"lower,omitempty"`
	Upper       string                `json:"upper,omitempty"`
	LowerStrict bool                  `json:"lowerStrict,omitempty"`
	UpperStrict bool                  `json:"upperStrict,omitempty"`
	Ordering    sorting.SortingOrder  `json:"ordering,omitempty"`
	Extraction  extraction.Extraction `json:"extractionFn,omitempty"`
}

func (b BoundFilter) Type() string {
	return BOUND.name()
}

type IntervalFilter struct {
	FilterType FilterType            `json:"type"`
	Dimension  string                `json:"dimension"`
	Intervals  []string              `json:"intervals"`
	Extraction extraction.Extraction `json:"extractionFn,omitempty"`
}

func (i IntervalFilter) Type() string {
	return INTERVAL.name()
}

type TrueFilter struct {
	FilterType FilterType `json:"type"`
}

func (t TrueFilter) Type() string {
	return TRUE.name()
}
