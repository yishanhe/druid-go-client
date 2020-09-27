package lookup

type Lookup interface {
	Type() string
}

type MapLookup struct {
	LookupType LookupType        `json:"type"`
	Map        map[string]string `json:"map"`
	IsOneToOne bool              `json:"isOneToOne,omitempty"`
}

func (m MapLookup) Type() string {
	return Map.Name()
}
