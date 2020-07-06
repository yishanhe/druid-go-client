package extraction

import "bytes"

type ExtractionType int

const (
	REGEX ExtractionType = iota
	PARTIAL
)

var extractionTypeStrings = []string{
	"regex", "partial",
}

func (e ExtractionType) name() string {
	return extractionTypeStrings[e]
}

func (e ExtractionType) ordinal() int {
	return int(e)
}

func (e ExtractionType) values() *[]string {
	return &extractionTypeStrings
}

func (e ExtractionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(e.name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
