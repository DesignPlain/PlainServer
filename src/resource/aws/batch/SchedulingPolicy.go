package batch

import types "DesignSphere_Server/src/resource/aws/types"

type SchedulingPolicy struct {
	//
	FairSharePolicy types.Batch_SchedulingPolicyFairSharePolicy `json:"fairSharePolicy,omitempty" yaml:"fairSharePolicy,omitempty"`

	// Specifies the name of the scheduling policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
