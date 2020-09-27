package lookup

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type LookupType int

const (
	Map LookupType = iota
)

var lookupTypeStrings = []string{"map"}

func (l LookupType) Name() string {
	return lookupTypeStrings[l]
}

func (l LookupType) Ordinal() int {
	return int(l)
}

func (l LookupType) Values() *[]string {
	return &lookupTypeStrings
}

func (l LookupType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(l.Name())
}
