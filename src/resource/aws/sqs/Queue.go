package sqs

type Queue struct {
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit string `json:"fifoThroughputLimit,omitempty" yaml:"fifoThroughputLimit,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId string `json:"kmsMasterKeyId,omitempty" yaml:"kmsMasterKeyId,omitempty"`

	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize int `json:"maxMessageSize,omitempty" yaml:"maxMessageSize,omitempty"`

	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds int `json:"messageRetentionSeconds,omitempty" yaml:"messageRetentionSeconds,omitempty"`

	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds int `json:"visibilityTimeoutSeconds,omitempty" yaml:"visibilityTimeoutSeconds,omitempty"`

	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue` (default).
	DeduplicationScope string `json:"deduplicationScope,omitempty" yaml:"deduplicationScope,omitempty"`

	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds int `json:"delaySeconds,omitempty" yaml:"delaySeconds,omitempty"`

	// The JSON policy to set up the Dead Letter Queue redrive permission, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html).
	RedriveAllowPolicy string `json:"redriveAllowPolicy,omitempty" yaml:"redriveAllowPolicy,omitempty"`

	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). --Note:-- when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy string `json:"redrivePolicy,omitempty" yaml:"redrivePolicy,omitempty"`

	// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html). The provider will only perform drift detection of its value when present in a configuration.
	SqsManagedSseEnabled bool `json:"sqsManagedSseEnabled,omitempty" yaml:"sqsManagedSseEnabled,omitempty"`

	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue bool `json:"fifoQueue,omitempty" yaml:"fifoQueue,omitempty"`

	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds int `json:"kmsDataKeyReusePeriodSeconds,omitempty" yaml:"kmsDataKeyReusePeriodSeconds,omitempty"`

	// The JSON policy for the SQS queue.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// A map of tags to assign to the queue. If configured with a provider `default_tags` configuration block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication bool `json:"contentBasedDeduplication,omitempty" yaml:"contentBasedDeduplication,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds int `json:"receiveWaitTimeSeconds,omitempty" yaml:"receiveWaitTimeSeconds,omitempty"`
}
