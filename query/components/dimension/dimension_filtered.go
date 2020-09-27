package dimension

type ListFilteredDimension struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Values        []string      `json:"values"`
	WhiteListed   bool          `json:"isWhitelist,omitempty"`
}

func (l ListFilteredDimension) Type() string {
	return ListFiltered.Name()
}

type RegexFilteredDimension struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Pattern       string        `json:"pattern"`
}

func (r RegexFilteredDimension) Type() string {
	return RegexFiltered.Name()
}

type PrefixFilteredDimension struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Prefix        string        `json:"prefix"`
}

func (p PrefixFilteredDimension) Type() string {
	return PrefixFiltered.Name()
}
