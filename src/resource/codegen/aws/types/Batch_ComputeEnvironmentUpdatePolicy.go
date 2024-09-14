package types

type Batch_ComputeEnvironmentUpdatePolicy struct {
	// Specifies the job timeout (in minutes) when the compute environment infrastructure is updated.
	JobExecutionTimeoutMinutes int `json:"jobExecutionTimeoutMinutes,omitempty" yaml:"jobExecutionTimeoutMinutes,omitempty"`

	// Specifies whether jobs are automatically terminated when the computer environment infrastructure is updated.
	TerminateJobsOnUpdate bool `json:"terminateJobsOnUpdate,omitempty" yaml:"terminateJobsOnUpdate,omitempty"`
}
