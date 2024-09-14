package types

type Batch_getJobDefinitionRetryStrategy struct {
	// Array of up to 5 objects that specify the conditions where jobs are retried or failed.
	EvaluateOnExits []Batch_getJobDefinitionRetryStrategyEvaluateOnExit `json:"evaluateOnExits,omitempty" yaml:"evaluateOnExits,omitempty"`

	// The number of times to move a job to the RUNNABLE status.
	Attempts int `json:"attempts,omitempty" yaml:"attempts,omitempty"`
}
