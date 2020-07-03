package granularity

type DurationGranularity struct {
	GranularityType GranularityType `json:"type"`
	Duration        int64           `json:"duration"`
	Origin          string          `json:"origin,omitempty"`
}

func (d DurationGranularity) Type() string {
	return DURATION.name()
}
