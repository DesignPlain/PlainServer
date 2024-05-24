package sqs

type QueuePolicy struct {
	// The JSON policy for the SQS queue.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `json:"queueUrl,omitempty" yaml:"queueUrl,omitempty"`
}
