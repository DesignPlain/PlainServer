package sqs

type RedriveAllowPolicy struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `json:"queueUrl,omitempty" yaml:"queueUrl,omitempty"`

	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy string `json:"redriveAllowPolicy,omitempty" yaml:"redriveAllowPolicy,omitempty"`
}
