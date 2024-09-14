package types

type Container_AwsNodePoolConfigTaint struct {
	// Value for the taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The taint effect. Possible values: EFFECT_UNSPECIFIED, NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Key for the taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
