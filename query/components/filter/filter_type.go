package filter

import "github.com/yishanhe/druid-go-client/pkg/enum"

type FilterType int

const (
	Selector FilterType = iota
	ColumnComparison
	Regex
	And
	Or
	Not
	Javascript
	Extraction
	Search
	In
	Like
	Bound
	Interval
	True
)

var filterTypeStrings = []string{
	"selector", "columnComparison", "regex", "and", "or", "not", "javascript", "extraction", "search", "in", "like", "bound", "interval", "true",
}

func (f FilterType) Name() string {
	return filterTypeStrings[f]
}

func (f FilterType) Ordinal() int {
	return int(f)
}

func (f FilterType) Values() *[]string {
	return &filterTypeStrings
}

func (f FilterType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(f.Name())
}
