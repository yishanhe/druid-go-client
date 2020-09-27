package enum

import "bytes"

type Enum interface {
	Name() string
	Ordinal() int
	Values() *[]string
}

func MarshalEnumJSON(name string) ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(name)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
