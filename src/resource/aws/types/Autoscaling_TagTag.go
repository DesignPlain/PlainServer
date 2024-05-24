package types

type Autoscaling_TagTag struct {
	// Tag name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Whether to propagate the tags to instances launched by the ASG.
	PropagateAtLaunch bool `json:"propagateAtLaunch,omitempty" yaml:"propagateAtLaunch,omitempty"`

	// Tag value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
