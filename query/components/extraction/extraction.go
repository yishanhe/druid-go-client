package extraction

type ExtractionFn interface {
	Type() string
}

type RegexExtractionFn struct {
	ExtractionType ExtractionType `json:"type"`
}

func (r RegexExtractionFn) Type() string {
	return REGEX.name()
}

