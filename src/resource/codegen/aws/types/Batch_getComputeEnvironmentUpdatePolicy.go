package types

type Batch_getComputeEnvironmentUpdatePolicy struct {
	//
	JobExecutionTimeoutMinutes int `json:"jobExecutionTimeoutMinutes,omitempty" yaml:"jobExecutionTimeoutMinutes,omitempty"`

	//
	TerminateJobsOnUpdate bool `json:"terminateJobsOnUpdate,omitempty" yaml:"terminateJobsOnUpdate,omitempty"`
}
