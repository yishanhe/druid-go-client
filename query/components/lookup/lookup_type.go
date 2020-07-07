package lookup

import "bytes"

type LookupType int

const (
	MAP LookupType = iota
)

var lookupTypeStrings = []string{"map"}

func (l LookupType) name() string {
	return lookupTypeStrings[l]
}

func (l LookupType) ordinal() int {
	return int(l)
}

func (l LookupType) values() *[]string {
	return &lookupTypeStrings
}

func (l LookupType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(l.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
