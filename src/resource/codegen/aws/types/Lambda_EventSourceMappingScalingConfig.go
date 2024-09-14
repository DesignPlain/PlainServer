package types

type Lambda_EventSourceMappingScalingConfig struct {
	// Limits the number of concurrent instances that the Amazon SQS event source can invoke. Must be greater than or equal to `2`. See [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency). You need to raise a [Service Quota Ticket](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) to increase the concurrency beyond 1000.
	MaximumConcurrency int `json:"maximumConcurrency,omitempty" yaml:"maximumConcurrency,omitempty"`
}
