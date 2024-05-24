package types

type Batch_JobDefinitionRetryStrategy struct {
	// The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
	Attempts int `json:"attempts,omitempty" yaml:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the `attempts` parameter must also be specified. You may specify up to 5 configuration blocks.
	EvaluateOnExits []Batch_JobDefinitionRetryStrategyEvaluateOnExit `json:"evaluateOnExits,omitempty" yaml:"evaluateOnExits,omitempty"`
}
