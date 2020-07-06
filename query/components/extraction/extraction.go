package extraction

type Extraction interface {
	Type() string
}

type RegexExtraction struct {
	ExtractionType ExtractionType `json:"type"`
}

func (r RegexExtraction) Type() string {
	return REGEX.name()
}
