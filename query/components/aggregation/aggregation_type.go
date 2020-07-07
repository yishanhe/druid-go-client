package aggregation


type AggregationType int

const (
	COUNT AggregationType = iota
	SUM
	MAX
	MIN
	FIRST
	LAST
	ANY
	JAVASCRIPT
	THETA_SKETCH
	FILTERED
)
