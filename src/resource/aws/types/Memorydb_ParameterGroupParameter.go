package types

type Memorydb_ParameterGroupParameter struct {
	// The value of the parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
