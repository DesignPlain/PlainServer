package batch

import types "DesignSphere_Server/src/resource/aws/types"

type JobQueue struct {
	// Specifies the name of the job queue.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The priority of the job queue. Job queues with a higher priority
	   are evaluated first when associated with the same compute environment.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn't specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can't remove the fair share scheduling policy.
	SchedulingPolicyArn string `json:"schedulingPolicyArn,omitempty" yaml:"schedulingPolicyArn,omitempty"`

	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Batch_JobQueueTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	/*
	   List of compute environment ARNs mapped to a job queue.
	   The position of the compute environments in the list will dictate the order.
	*/
	ComputeEnvironments []string `json:"computeEnvironments,omitempty" yaml:"computeEnvironments,omitempty"`
}
