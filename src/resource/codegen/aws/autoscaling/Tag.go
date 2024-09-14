package autoscaling

import types "libds/aws/types"

type Tag struct {
	// Name of the Autoscaling Group to apply the tag to.
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`

	// Tag to create. The `tag` block is documented below.
	Tag types.Autoscaling_TagTag `json:"tag,omitempty" yaml:"tag,omitempty"`
}
