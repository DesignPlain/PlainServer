package types

type Batch_getComputeEnvironmentUpdatePolicy struct {
	//
	TerminateJobsOnUpdate bool `json:"terminateJobsOnUpdate,omitempty" yaml:"terminateJobsOnUpdate,omitempty"`

	//
	JobExecutionTimeoutMinutes int `json:"jobExecutionTimeoutMinutes,omitempty" yaml:"jobExecutionTimeoutMinutes,omitempty"`
}
