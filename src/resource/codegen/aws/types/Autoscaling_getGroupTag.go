package types

type Autoscaling_getGroupTag struct {
	// Key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Whether the tag is propagated to Amazon EC2 instances launched via this ASG.
	PropagateAtLaunch bool `json:"propagateAtLaunch,omitempty" yaml:"propagateAtLaunch,omitempty"`

	// Value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
