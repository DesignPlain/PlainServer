package types

type Pipes_PipeSourceParametersSqsQueueParameters struct {
	// The maximum number of records to include in each batch. Maximum value of 10000.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// The maximum length of a time to wait for events. Maximum value of 300.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`
}
