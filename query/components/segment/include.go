package segment

type Include interface {
	Type() string
}

type AllInclude struct {
	IncludeType IncludeType `json:"type"`
}

func (a AllInclude) Type() string {
	return All.Name()
}

type NoneInclude struct {
	IncludeType IncludeType `json:"type"`
}

func (n NoneInclude) Type() string {
	return None.Name()
}

type ListInclude struct {
	IncludeType IncludeType `json:"type"`
	Columns     []string    `json:"columns"`
}

func (l ListInclude) Type() string {
	return List.Name()
}
