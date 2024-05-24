package kinesis

import types "DesignSphere_Server/src/resource/aws/types"

type Stream struct {
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics []string `json:"shardLevelMetrics,omitempty" yaml:"shardLevelMetrics,omitempty"`

	// Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
	StreamModeDetails types.Kinesis_StreamStreamModeDetails `json:"streamModeDetails,omitempty" yaml:"streamModeDetails,omitempty"`

	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion bool `json:"enforceConsumerDeletion,omitempty" yaml:"enforceConsumerDeletion,omitempty"`

	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	/*
	   The number of shards that the stream will use. If the `stream_mode` is `PROVISIONED`, this field is required.
	   Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	*/
	ShardCount int `json:"shardCount,omitempty" yaml:"shardCount,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
