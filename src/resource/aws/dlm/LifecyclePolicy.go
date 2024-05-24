package dlm

import types "DesignSphere_Server/src/resource/aws/types"

type LifecyclePolicy struct {
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description for the DLM lifecycle policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// See the `policy_details` configuration block. Max of 1.
	PolicyDetails types.Dlm_LifecyclePolicyPolicyDetails `json:"policyDetails,omitempty" yaml:"policyDetails,omitempty"`
}
