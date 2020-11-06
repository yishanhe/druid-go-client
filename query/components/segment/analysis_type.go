package segment

import "github.com/yishanhe/druid-go-client/pkg/enum"

type AnalysisType int

const (
	Cardinality AnalysisType = iota
	Minmax
	Size
	Interval
	TimestampSpec
	QueryGranularity
	Aggregators
	Rollup
)

var analysisTypeStrings = []string{"cardinality", "minmax", "size", "interval", "timestampSpec", "queryGranularity",
	"aggregators", "rollup"}

func (a AnalysisType) Name() string {
	return analysisTypeStrings[a]
}

func (a AnalysisType) Ordinal() int {
	return int(a)
}

func (a AnalysisType) Values() *[]string {
	return &analysisTypeStrings
}

func (a AnalysisType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnumJSON(a.Name())
}
