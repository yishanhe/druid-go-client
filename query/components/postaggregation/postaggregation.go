package postaggregation

type PostAggregator interface {
	Type() string
}

type ArithmeticPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name string `json:"name"`
	Function string `json:"function"`
	Fields PostAggregator `json:"fields"`
	Ordering string `json:"ordering"`
}

func (a ArithmeticPostAggregation) Type() string {
	return ARITHMETIC.name()
}

// FieldAccessPostAggregation can be hyperUniqueCardinality, fieldAccess, finalizingFieldAccess
type FieldAccessPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name string `json:"name"`
	FieldName string `json:"fieldName"`
}

func (f FieldAccessPostAggregation) Type() string {
	return f.PostAggregationType.name()
}

type ConstantPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name string `json:"name"`
	Value string `json:"value"`
}

func (c ConstantPostAggregation) Type() string {
	return CONSTANT.name()
}

type GreatestPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name string `json:"name"`
	Fields PostAggregator `json:"fields"`
}

func (g GreatestPostAggregation) Type() string {
	return g.PostAggregationType.name()
}

type LeastPostAggregation struct {
	PostAggregationType PostAggregationType `json:"type"`
	Name string `json:"name"`
	Fields PostAggregator `json:"fields"`
}

func (l LeastPostAggregation) Type() string {
	return l.PostAggregationType.name()
}





