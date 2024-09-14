package types

type Pipes_PipeSourceParametersDynamodbStreamParameters struct {
	// Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source. Maximum value of 10,000.
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" yaml:"maximumRetryAttempts,omitempty"`

	// Define how to handle item process failures. AUTOMATIC_BISECT halves each batch and retry each half until all the records are processed or there is one failed message left in the batch. Valid values: AUTOMATIC_BISECT.
	OnPartialBatchItemFailure string `json:"onPartialBatchItemFailure,omitempty" yaml:"onPartialBatchItemFailure,omitempty"`

	// The number of batches to process concurrently from each shard. The default value is 1. Maximum value of 10.
	ParallelizationFactor int `json:"parallelizationFactor,omitempty" yaml:"parallelizationFactor,omitempty"`

	// The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
	StartingPosition string `json:"startingPosition,omitempty" yaml:"startingPosition,omitempty"`

	// The maximum number of records to include in each batch. Maximum value of 10000.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// Define the target queue to send dead-letter queue events to. Detailed below.
	DeadLetterConfig Pipes_PipeSourceParametersDynamodbStreamParametersDeadLetterConfig `json:"deadLetterConfig,omitempty" yaml:"deadLetterConfig,omitempty"`

	// The maximum length of a time to wait for events. Maximum value of 300.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`

	// Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records. Maximum value of 604,800.
	MaximumRecordAgeInSeconds int `json:"maximumRecordAgeInSeconds,omitempty" yaml:"maximumRecordAgeInSeconds,omitempty"`
}
