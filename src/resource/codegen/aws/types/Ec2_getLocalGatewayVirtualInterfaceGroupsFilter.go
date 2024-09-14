package types

type Ec2_getLocalGatewayVirtualInterfaceGroupsFilter struct {
	// Name of the filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
