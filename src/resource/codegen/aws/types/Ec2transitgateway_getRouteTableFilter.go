package types

type Ec2transitgateway_getRouteTableFilter struct {
	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of the filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
