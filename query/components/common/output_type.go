package common

import (
	"github.com/yishanhe/druid-go-client/pkg/enum"
)

type OutputType int

const (
	String OutputType = iota
	Long
	Float
)

var outputTypeStrings = []string{
	"STRING", "LONG", "FLOAT",
}

func (o OutputType) Name() string {
	return outputTypeStrings[o]
}

func (o OutputType) Ordinal() int {
	return int(o)
}

func (o OutputType) Values() *[]string {
	return &outputTypeStrings
}

func (o OutputType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(o.Name())
}
