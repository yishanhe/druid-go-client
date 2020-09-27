package search

type Search interface {
	Type() string
}

type InsensitiveContainsSearch struct {
	SearchType SearchType `json:"type"`
	Value      string     `json:"value"`
}

func (i InsensitiveContainsSearch) Type() string {
	return InsensitiveContains.Name()
}

type FragmentSearch struct {
	SearchType    SearchType `json:"type"`
	CaseSensitive bool       `json:"case_sensitive"`
	Values        []string   `json:"values"`
}

func (f FragmentSearch) Type() string {
	return Fragment.Name()
}

type ContainsSearch struct {
	SearchType    SearchType `json:"type"`
	CaseSensitive bool       `json:"case_sensitive"`
	Value         string     `json:"value"`
}

func (c ContainsSearch) Type() string {
	return Contains.Name()
}

type RegexSearch struct {
	SearchType SearchType `json:"type"`
	Pattern    string     `json:"pattern"`
}

func (r RegexSearch) Type() string {
	return Regex.Name()
}
