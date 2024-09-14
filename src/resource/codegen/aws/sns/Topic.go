package sns

type Topic struct {
	// The IAM role permitted to receive success feedback for this topic
	FirehoseSuccessFeedbackRoleArn string `json:"firehoseSuccessFeedbackRoleArn,omitempty" yaml:"firehoseSuccessFeedbackRoleArn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	HttpSuccessFeedbackRoleArn string `json:"httpSuccessFeedbackRoleArn,omitempty" yaml:"httpSuccessFeedbackRoleArn,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
	KmsMasterKeyId string `json:"kmsMasterKeyId,omitempty" yaml:"kmsMasterKeyId,omitempty"`

	// IAM role for failure feedback
	ApplicationFailureFeedbackRoleArn string `json:"applicationFailureFeedbackRoleArn,omitempty" yaml:"applicationFailureFeedbackRoleArn,omitempty"`

	// If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	SignatureVersion int `json:"signatureVersion,omitempty" yaml:"signatureVersion,omitempty"`

	// IAM role for failure feedback
	HttpFailureFeedbackRoleArn string `json:"httpFailureFeedbackRoleArn,omitempty" yaml:"httpFailureFeedbackRoleArn,omitempty"`

	// The fully-formed AWS policy as JSON.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// IAM role for failure feedback
	SqsFailureFeedbackRoleArn string `json:"sqsFailureFeedbackRoleArn,omitempty" yaml:"sqsFailureFeedbackRoleArn,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	ApplicationSuccessFeedbackRoleArn string `json:"applicationSuccessFeedbackRoleArn,omitempty" yaml:"applicationSuccessFeedbackRoleArn,omitempty"`

	// Percentage of success to sample
	ApplicationSuccessFeedbackSampleRate int `json:"applicationSuccessFeedbackSampleRate,omitempty" yaml:"applicationSuccessFeedbackSampleRate,omitempty"`

	// The message archive policy for FIFO topics. More details in the [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/message-archiving-and-replay-topic-owner.html).
	ArchivePolicy string `json:"archivePolicy,omitempty" yaml:"archivePolicy,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	LambdaSuccessFeedbackRoleArn string `json:"lambdaSuccessFeedbackRoleArn,omitempty" yaml:"lambdaSuccessFeedbackRoleArn,omitempty"`

	// The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Percentage of success to sample
	SqsSuccessFeedbackSampleRate int `json:"sqsSuccessFeedbackSampleRate,omitempty" yaml:"sqsSuccessFeedbackSampleRate,omitempty"`

	// Percentage of success to sample
	LambdaSuccessFeedbackSampleRate int `json:"lambdaSuccessFeedbackSampleRate,omitempty" yaml:"lambdaSuccessFeedbackSampleRate,omitempty"`

	// The display name for the topic
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic. FIFO topics can't deliver messages to customer managed endpoints, such as email addresses, mobile apps, SMS, or HTTP(S) endpoints. These endpoint types aren't guaranteed to preserve strict message ordering. Default is `false`.
	FifoTopic bool `json:"fifoTopic,omitempty" yaml:"fifoTopic,omitempty"`

	// Percentage of success to sample
	HttpSuccessFeedbackSampleRate int `json:"httpSuccessFeedbackSampleRate,omitempty" yaml:"httpSuccessFeedbackSampleRate,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	SqsSuccessFeedbackRoleArn string `json:"sqsSuccessFeedbackRoleArn,omitempty" yaml:"sqsSuccessFeedbackRoleArn,omitempty"`

	// Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
	ContentBasedDeduplication bool `json:"contentBasedDeduplication,omitempty" yaml:"contentBasedDeduplication,omitempty"`

	// IAM role for failure feedback
	FirehoseFailureFeedbackRoleArn string `json:"firehoseFailureFeedbackRoleArn,omitempty" yaml:"firehoseFailureFeedbackRoleArn,omitempty"`

	// Percentage of success to sample
	FirehoseSuccessFeedbackSampleRate int `json:"firehoseSuccessFeedbackSampleRate,omitempty" yaml:"firehoseSuccessFeedbackSampleRate,omitempty"`

	// IAM role for failure feedback
	LambdaFailureFeedbackRoleArn string `json:"lambdaFailureFeedbackRoleArn,omitempty" yaml:"lambdaFailureFeedbackRoleArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Tracing mode of an Amazon SNS topic. Valid values: `"PassThrough"`, `"Active"`.
	TracingConfig string `json:"tracingConfig,omitempty" yaml:"tracingConfig,omitempty"`

	// The SNS delivery policy. More details in the [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html).
	DeliveryPolicy string `json:"deliveryPolicy,omitempty" yaml:"deliveryPolicy,omitempty"`
}
