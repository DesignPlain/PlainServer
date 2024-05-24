package sqs

type RedrivePolicy struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `json:"queueUrl,omitempty" yaml:"queueUrl,omitempty"`

	// The JSON redrive policy for the SQS queue. Accepts two key/val pairs: `deadLetterTargetArn` and `maxReceiveCount`. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedrivePolicy string `json:"redrivePolicy,omitempty" yaml:"redrivePolicy,omitempty"`
}
