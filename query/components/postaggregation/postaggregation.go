package postaggregation

type PostAggregator interface {
	Type() string
}

type ArithmeticPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name,omitempty"`
	Function            string              `json:"fn"`
	Fields              []PostAggregator    `json:"fields"`
	Ordering            string              `json:"ordering,omitempty"`
}

func (a ArithmeticPostAggregation) Type() string {
	return Arithmetic.Name()
}

// FieldAccessPostAggregation can be hyperUniqueCardinality, fieldAccess, finalizingFieldAccess
type FieldAccessPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name,omitempty"`
	FieldName           string              `json:"fieldName"`
}

func (f FieldAccessPostAggregation) Type() string {
	return f.PostAggregationType.Name()
}

type JavascriptPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name,omitempty"`
	FieldNames          []string            `json:"fieldNames"`
	Function            string              `json:"function"`
}

func (j JavascriptPostAggregation) Type() string {
	return j.PostAggregationType.Name()
}

type ConstantPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name"`
	Value               string              `json:"value"`
}

func (c ConstantPostAggregation) Type() string {
	return Constant.Name()
}

// DoubleGreatest, LongGreatest
type GreatestPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name"`
	Fields              []PostAggregator    `json:"fields"`
}

func (g GreatestPostAggregation) Type() string {
	return g.PostAggregationType.Name()
}

// DoubleLeast, LongLeast
type LeastPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name                string              `json:"name"`
	Fields              []PostAggregator    `json:"fields"`
}

func (l LeastPostAggregation) Type() string {
	return l.PostAggregationType.Name()
}
