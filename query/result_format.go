package query

import "bytes"

type ResultFormat int

const (
	LIST ResultFormat = iota
	COMPACTED_LIST
)

var resultFormatStrings = []string{"list", "compactedList"}

func (r ResultFormat) name() string {
	return resultFormatStrings[r]
}

func (r ResultFormat) ordinal() int {
	return int(r)
}

func (r ResultFormat) values() *[]string {
	return &resultFormatStrings
}

func (r ResultFormat) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil

}
