package types

type Codedeploy_DeploymentGroupEc2TagFilter struct {
	// The type of the tag filter, either `KEY_ONLY`, `VALUE_ONLY`, or `KEY_AND_VALUE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The value of the tag filter.

	   Multiple occurrences of `ec2_tag_filter` are allowed, where any instance that matches to at least one of the tag filters is selected.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The key of the tag filter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
