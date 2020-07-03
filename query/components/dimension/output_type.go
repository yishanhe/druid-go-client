package dimension

import "bytes"

type OutputType int

const (
	STRING OutputType = iota
	LONG
	FLOAT
)

var outputTypeStrings = []string{
	"STRING", "LONG", "FLOAT",
}

func (o OutputType) name() string {
	return outputTypeStrings[o]
}

func (o OutputType) ordinal() int {
	return int(o)
}

func (o OutputType) values() *[]string {
	return &outputTypeStrings
}

func (o OutputType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(o.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
