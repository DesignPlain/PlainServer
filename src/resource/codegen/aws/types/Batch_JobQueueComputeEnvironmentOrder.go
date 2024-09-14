package types

type Batch_JobQueueComputeEnvironmentOrder struct {
	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironment string `json:"computeEnvironment,omitempty" yaml:"computeEnvironment,omitempty"`

	// The order of the compute environment. Compute environments are tried in ascending order. For example, if two compute environments are associated with a job queue, the compute environment with a lower order integer value is tried for job placement first.
	Order int `json:"order,omitempty" yaml:"order,omitempty"`
}
