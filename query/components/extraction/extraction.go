package extraction

import (
	"github.com/yishanhe/druid-go-client/query/components/granularity"
	"github.com/yishanhe/druid-go-client/query/components/lookup"
	"github.com/yishanhe/druid-go-client/query/components/search"
)

type Extraction interface {
	Type() string
}

type RegexExtraction struct {
	ExtractionType          ExtractionType `json:"type"`
	Expression              string         `json:"expr"`
	Index                   int            `json:"index,omitempty"`
	ReplaceMissingValue     bool           `json:"replaceMissingValue,omitempty"`
	ReplaceMissingValueWith string         `json:"replaceMissingValueWith,omitempty"`
}

func (r RegexExtraction) Type() string {
	return Regex.Name()
}

type PartialExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Expression     string         `json:"expr"`
}

func (p PartialExtraction) Type() string {
	return Partial.Name()
}

type SearchQueryExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	SearchQuery    search.Search  `json:"query"`
}

func (s SearchQueryExtraction) Type() string {
	return SearchQuery.Name()
}

type SubstringExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Index          int            `json:"index"`
	Length         int            `json:"length,omitempty"`
}

func (s SubstringExtraction) Type() string {
	return Substring.Name()
}

type StrlenExtraction struct {
	ExtractionType ExtractionType `json:"type"`
}

func (s StrlenExtraction) Type() string {
	return Strlen.Name()
}

type TimeFormatExtraction struct {
	ExtractionType ExtractionType          `json:"type"`
	Format         string                  `json:"format,omitempty"`
	TimeZone       string                  `json:"timeZone,omitempty"`
	Locale         string                  `json:"locale,omitempty"`
	Granularity    granularity.Granularity `json:"granularity,omitempty"`
	AsMillis       bool                    `json:"asMillis,omitempty"`
}

func (t TimeFormatExtraction) Type() string {
	return TimeFormat.Name()
}

type TimeExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	TimeFormat     string         `json:"timeFormat"`
	ResultFormat   string         `json:"resultFormat"`
	Joda           bool           `json:"joda,omitempty"`
}

func (t TimeExtraction) Type() string {
	return Time.Name()
}

type JavascriptExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Function       string         `json:"function"`
	Injective      bool           `json:"injective,omitempty"`
}

func (j JavascriptExtraction) Type() string {
	return Javascript.Name()
}

type RegisteredLookupExtraction struct {
	ExtractionType     ExtractionType `json:"type"`
	Lookup             string         `json:"lookup"`
	RetainMissingValue bool           `json:"retainMissingValue,omitempty"`
}

func (r RegisteredLookupExtraction) Type() string {
	return RegisteredLookup.Name()
}

type LookupExtraction struct {
	ExtractionType          ExtractionType `json:"type"`
	Lookup                  lookup.Lookup  `json:"lookup"`
	RetainMissingValue      bool           `json:"retainMissingValue,omitempty"`
	Injective               bool           `json:"injective,omitempty"`
	ReplaceMissingValueWith string         `json:"replaceMissingValueWith,omitempty"`
}

func (l LookupExtraction) Type() string {
	return Lookup.Name()
}

type CascadeExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Extractions    []Extraction   `json:"extractionFns"`
}

func (c CascadeExtraction) Type() string {
	return Cascade.Name()
}

type StringFormatExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Format         string         `json:"format"`
	NullHandling   string         `json:"nullHandling,omitempty"`
}

func (s StringFormatExtraction) Type() string {
	return StringFormat.Name()
}

type UpperExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Locale         string         `json:"locale,omitempty"`
}

func (u UpperExtraction) Type() string {
	return Upper.Name()
}

type LowerExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Locale         string         `json:"locale,omitempty"`
}

func (l LowerExtraction) Type() string {
	return Lower.Name()
}

type BucketExtraction struct {
	ExtractionType ExtractionType `json:"type"`
	Size           int64          `json:"size,omitempty"`
	Offset         int64          `json:"offset,omitempty"`
}

func (b BucketExtraction) Type() string {
	return Bucket.Name()
}
