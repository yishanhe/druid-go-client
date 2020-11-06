package segment

import "github.com/yishanhe/druid-go-client/pkg/enum"

type IncludeType int

const (
	All IncludeType = iota
	None
	List
)

var includeTypeStrings = []string{"all", "none", "list"}

func (i IncludeType) Name() string {
	return includeTypeStrings[i]
}

func (i IncludeType) Ordinal() int {
	return int(i)
}

func (i IncludeType) Values() *[]string {
	return &includeTypeStrings
}

func (i IncludeType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(i.Name())
}
