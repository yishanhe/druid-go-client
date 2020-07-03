package dimension

type ListFilteredDimensionSpec struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Values        []string      `json:"values"`
	WhiteListed   bool          `json:"isWhitelist,omitempty"`
}

func (l ListFilteredDimensionSpec) Type() string {
	return LIST_FILTERED.name()
}

type RegexFilteredDimension struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Pattern       string        `json:"pattern"`
}

func (r RegexFilteredDimension) Type() string {
	return REGEX_FILTERED.name()
}

type PrefixFilteredDimension struct {
	DimensionType DimensionType `json:"type"`
	Delegate      Dimension     `json:"delegate"`
	Prefix        string        `json:"prefix"`
}

func (p PrefixFilteredDimension) Type() string {
	return PREFIX_FILTERED.name()
}
