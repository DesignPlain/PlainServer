package types

type Memorydb_getParameterGroupParameter struct {
	// Name of the parameter group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
