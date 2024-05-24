package types

type Ec2_getVpcDhcpOptionsFilter struct {
	// Name of the field to filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values for filtering.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
