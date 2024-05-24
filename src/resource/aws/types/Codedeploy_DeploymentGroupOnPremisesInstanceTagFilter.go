package types

type Codedeploy_DeploymentGroupOnPremisesInstanceTagFilter struct {
	// The key of the tag filter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The type of the tag filter, either `KEY_ONLY`, `VALUE_ONLY`, or `KEY_AND_VALUE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The value of the tag filter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
