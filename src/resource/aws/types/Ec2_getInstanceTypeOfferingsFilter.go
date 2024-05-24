package types

type Ec2_getInstanceTypeOfferingsFilter struct {
	// Name of the filter. The `location` filter depends on the top-level `location_type` argument and if not specified, defaults to the current region.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
