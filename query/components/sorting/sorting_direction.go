package sorting

import "github.com/yishanhe/druid-go-client/pkg/enum"

type Direction int

const (
	Ascending Direction = iota
	Descending
)

var directionStrings = []string{
	"ascending", "descending",
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
