package types

type Autoscaling_getAmiIdsFilter struct {
	// Name of the DescribeAutoScalingGroup filter. The recommended values are: `tag-key`, `tag-value`, and `tag:<tag name>`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
