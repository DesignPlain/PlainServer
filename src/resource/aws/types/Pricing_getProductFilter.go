package types

type Pricing_getProductFilter struct {
	// Product attribute value that you want to filter on.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Product attribute name that you want to filter on.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`
}
