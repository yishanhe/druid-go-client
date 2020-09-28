package limit

import "github.com/yishanhe/druid-go-client/pkg/enum"

type LimitType int

const (
	Default LimitType = iota
)

var limitTypeStrings = []string{
	"default",
}

func (l LimitType) Name() string {
	return limitTypeStrings[l]
}

func (l LimitType) Ordinal() int {
	return int(l)
}

func (l LimitType) Values() *[]string {
	return &limitTypeStrings
}

func (l LimitType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(l.Name())
}
