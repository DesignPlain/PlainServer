package types

type Ec2_getVpcIpamPoolFilter struct {
	// The name of the filter. Filter names are case-sensitive.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The filter values. Filter values are case-sensitive.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
