package autoscaling

import types "DesignSphere_Server/src/resource/aws/types"

type TrafficSourceAttachment struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`

	// The unique identifiers of a traffic sources.
	TrafficSource types.Autoscaling_TrafficSourceAttachmentTrafficSource `json:"trafficSource,omitempty" yaml:"trafficSource,omitempty"`
}
