package having

import "github.com/yishanhe/druid-go-client/pkg/enum"

type HavingType int

const (
	Filter HavingType = iota
	EqualTo
	GreaterThan
	LessThan
	DimSelector
	And
	Or
	Not
)

var havingTypeStrings = []string{
	"filter", "equalTo", "greaterThan", "lessThan", "dimSelector", "and", "or", "not",
}

func (h HavingType) Name() string {
	return havingTypeStrings[h]
}

func (h HavingType) Ordinal() int {
	return int(h)
}

func (h HavingType) Values() *[]string {
	return &havingTypeStrings
}

func (h HavingType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(h.Name())
}
