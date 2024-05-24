package emr

import types "DesignSphere_Server/src/resource/aws/types"

type ManagedScalingPolicy struct {
	// ID of the EMR cluster
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// Configuration block with compute limit settings. Described below.
	ComputeLimits []types.Emr_ManagedScalingPolicyComputeLimit `json:"computeLimits,omitempty" yaml:"computeLimits,omitempty"`
}
