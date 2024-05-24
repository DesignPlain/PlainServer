package batch

import types "DesignSphere_Server/src/resource/aws/types"

type ComputeEnvironment struct {
	// The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, the provider will assign a random, unique name.
	ComputeEnvironmentName string `json:"computeEnvironmentName,omitempty" yaml:"computeEnvironmentName,omitempty"`

	// Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
	ComputeEnvironmentNamePrefix string `json:"computeEnvironmentNamePrefix,omitempty" yaml:"computeEnvironmentNamePrefix,omitempty"`

	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of the compute environment. Valid items are `MANAGED` or `UNMANAGED`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies the infrastructure update policy for the compute environment. See details below.
	UpdatePolicy types.Batch_ComputeEnvironmentUpdatePolicy `json:"updatePolicy,omitempty" yaml:"updatePolicy,omitempty"`

	// Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
	ComputeResources types.Batch_ComputeEnvironmentComputeResources `json:"computeResources,omitempty" yaml:"computeResources,omitempty"`

	// Details for the Amazon EKS cluster that supports the compute environment. See details below.
	EksConfiguration types.Batch_ComputeEnvironmentEksConfiguration `json:"eksConfiguration,omitempty" yaml:"eksConfiguration,omitempty"`

	// The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
