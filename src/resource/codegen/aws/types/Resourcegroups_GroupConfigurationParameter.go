package types

type Resourcegroups_GroupConfigurationParameter struct {
	// The name of the group configuration parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value or values to be used for the specified parameter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
