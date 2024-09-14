package types

type Resourcegroups_GroupConfiguration struct {
	// A collection of parameters for this group configuration item. See below for details.
	Parameters []Resourcegroups_GroupConfigurationParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Specifies the type of group configuration item.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
