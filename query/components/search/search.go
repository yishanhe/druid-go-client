package search

type Search interface {
	Type() string
}

type InsensitiveContainsSearch struct {
	SearchType SearchType `json:"type"`
	Value      string     `json:"value"`
}

func (i InsensitiveContainsSearch) Type() string {
	return INSENSITIVE_CONTAINS.name()
}

type FragmentSearch struct {
	SearchType    SearchType `json:"type"`
	CaseSensitive bool       `json:"case_sensitive"`
	Values        []string   `json:"values"`
}

func (f FragmentSearch) Type() string {
	return FRAGMENT.name()
}

type ContainsSearch struct {
	SearchType    SearchType `json:"type"`
	CaseSensitive bool       `json:"case_sensitive"`
	Value         string     `json:"value"`
}

func (c ContainsSearch) Type() string {
	return CONTAINS.name()
}

type RegexSearch struct {
	SearchType SearchType `json:"type"`
	Pattern    string     `json:"pattern"`
}

func (r RegexSearch) Type() string {
	return REGEX.name()
}
