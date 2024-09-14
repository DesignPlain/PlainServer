package types

type Cloudwatch_EventTargetBatchTarget struct {
	// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
	JobDefinition string `json:"jobDefinition,omitempty" yaml:"jobDefinition,omitempty"`

	// The name to use for this execution of the job, if the target is an AWS Batch job.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
	ArraySize int `json:"arraySize,omitempty" yaml:"arraySize,omitempty"`

	// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
	JobAttempts int `json:"jobAttempts,omitempty" yaml:"jobAttempts,omitempty"`
}
