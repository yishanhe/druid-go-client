package granularity

type PeriodGranularity struct {
	GranularityType GranularityType `json:"type"`
	Period          string          `json:"period"`
	TimeZone        string          `json:"timeZone"`
	Origin          string          `json:"origin,omitempty"`
}

func (p PeriodGranularity) Type() string {
	return Period.Name()
}
