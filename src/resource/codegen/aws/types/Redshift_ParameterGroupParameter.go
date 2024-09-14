package types

type Redshift_ParameterGroupParameter struct {
	// The value of the Redshift parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the Redshift parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
