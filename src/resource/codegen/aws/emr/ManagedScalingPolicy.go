package emr

import types "libds/aws/types"

type ManagedScalingPolicy struct {
	// ID of the EMR cluster
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// Configuration block with compute limit settings. Described below.
	ComputeLimits []types.Emr_ManagedScalingPolicyComputeLimit `json:"computeLimits,omitempty" yaml:"computeLimits,omitempty"`
}
