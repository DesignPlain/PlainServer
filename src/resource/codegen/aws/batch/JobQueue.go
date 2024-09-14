package batch

import types "libds/aws/types"

type JobQueue struct {
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Batch_JobQueueTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// (Optional) This parameter is deprecated, please use `compute_environment_order` instead. List of compute environment ARNs mapped to a job queue. The position of the compute environments in the list will dictate the order. When importing a AWS Batch Job Queue, the parameter `compute_environments` will always be used over `compute_environment_order`. Please adjust your HCL accordingly.
	ComputeEnvironments []string `json:"computeEnvironments,omitempty" yaml:"computeEnvironments,omitempty"`

	/*
	   The priority of the job queue. Job queues with a higher priority
	   are evaluated first when associated with the same compute environment.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn't specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can't remove the fair share scheduling policy.
	SchedulingPolicyArn string `json:"schedulingPolicyArn,omitempty" yaml:"schedulingPolicyArn,omitempty"`

	// The set of compute environments mapped to a job queue and their order relative to each other. The job scheduler uses this parameter to determine which compute environment runs a specific job. Compute environments must be in the VALID state before you can associate them with a job queue. You can associate up to three compute environments with a job queue.
	ComputeEnvironmentOrders []types.Batch_JobQueueComputeEnvironmentOrder `json:"computeEnvironmentOrders,omitempty" yaml:"computeEnvironmentOrders,omitempty"`

	// The set of job state time limit actions mapped to a job queue. Specifies an action that AWS Batch will take after the job has remained at the head of the queue in the specified state for longer than the specified time.
	JobStateTimeLimitActions []types.Batch_JobQueueJobStateTimeLimitAction `json:"jobStateTimeLimitActions,omitempty" yaml:"jobStateTimeLimitActions,omitempty"`

	// Specifies the name of the job queue.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
