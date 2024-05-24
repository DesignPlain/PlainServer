package types

type Autoscaling_getAmiIdsFilter struct {
	// Value of the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of the DescribeAutoScalingGroup filter. The recommended values are: `tag-key`, `tag-value`, and `tag:<tag name>`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
