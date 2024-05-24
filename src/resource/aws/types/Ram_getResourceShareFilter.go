package types

type Ram_getResourceShareFilter struct {
	// Name of the tag key to filter on.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the tag key.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
