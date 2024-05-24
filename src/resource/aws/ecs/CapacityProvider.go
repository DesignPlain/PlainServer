package ecs

import types "DesignSphere_Server/src/resource/aws/types"

type CapacityProvider struct {
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider types.Ecs_CapacityProviderAutoScalingGroupProvider `json:"autoScalingGroupProvider,omitempty" yaml:"autoScalingGroupProvider,omitempty"`

	// Name of the capacity provider.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
