package sorting

import "github.com/yishanhe/druid-go-client/pkg/enum"

type Direction int

const (
	None Direction = iota
	Ascending
	Descending
)

var directionStrings = []string{
	"none", "ascending", "descending",
}

func (d Direction) Name() string {
	return directionStrings[d]
}

func (d Direction) Ordinal() int {
	return int(d)
}

func (d Direction) Values() *[]string {
	return &directionStrings
}

func (d Direction) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(d.Name())
}
