package types

type Pipes_PipeTargetParametersBatchJobParametersRetryStrategy struct {
	// The number of times to move a job to the RUNNABLE status. If the value of attempts is greater than one, the job is retried on failure the same number of attempts as the value. Maximum value of 10.
	Attempts int `json:"attempts,omitempty" yaml:"attempts,omitempty"`
}
