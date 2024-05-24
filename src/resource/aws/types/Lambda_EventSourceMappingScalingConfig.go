package types

type Lambda_EventSourceMappingScalingConfig struct {
	// Limits the number of concurrent instances that the Amazon SQS event source can invoke. Must be between `2` and `1000`. See [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).
	MaximumConcurrency int `json:"maximumConcurrency,omitempty" yaml:"maximumConcurrency,omitempty"`
}
