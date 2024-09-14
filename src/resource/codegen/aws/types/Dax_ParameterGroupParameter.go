package types

type Dax_ParameterGroupParameter struct {
	// The name of the parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value for the parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
