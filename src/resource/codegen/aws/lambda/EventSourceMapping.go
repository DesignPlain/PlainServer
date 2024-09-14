package lambda

import types "libds/aws/types"

type EventSourceMapping struct {
	// - (Optional) For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include `source_access_configuration`. Detailed below.
	SelfManagedEventSource types.Lambda_EventSourceMappingSelfManagedEventSource `json:"selfManagedEventSource,omitempty" yaml:"selfManagedEventSource,omitempty"`

	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB, MSK or Self Managed Apache Kafka. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition string `json:"startingPosition,omitempty" yaml:"startingPosition,omitempty"`

	// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
	Topics []string `json:"topics,omitempty" yaml:"topics,omitempty"`

	// - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	BisectBatchOnFunctionError bool `json:"bisectBatchOnFunctionError,omitempty" yaml:"bisectBatchOnFunctionError,omitempty"`

	// The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker, MSK cluster or DocumentDB change stream.  It is incompatible with a Self Managed Kafka source.
	EventSourceArn string `json:"eventSourceArn,omitempty" yaml:"eventSourceArn,omitempty"`

	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
	MaximumRecordAgeInSeconds int `json:"maximumRecordAgeInSeconds,omitempty" yaml:"maximumRecordAgeInSeconds,omitempty"`

	// Scaling configuration of the event source. Only available for SQS queues. Detailed below.
	ScalingConfig types.Lambda_EventSourceMappingScalingConfig `json:"scalingConfig,omitempty" yaml:"scalingConfig,omitempty"`

	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `starting_position` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	StartingPositionTimestamp string `json:"startingPositionTimestamp,omitempty" yaml:"startingPositionTimestamp,omitempty"`

	// The duration in seconds of a processing window for [AWS Lambda streaming analytics](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-windows). The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
	TumblingWindowInSeconds int `json:"tumblingWindowInSeconds,omitempty" yaml:"tumblingWindowInSeconds,omitempty"`

	// Additional configuration block for Amazon Managed Kafka sources. Incompatible with "self_managed_event_source" and "self_managed_kafka_event_source_config". Detailed below.
	AmazonManagedKafkaEventSourceConfig types.Lambda_EventSourceMappingAmazonManagedKafkaEventSourceConfig `json:"amazonManagedKafkaEventSourceConfig,omitempty" yaml:"amazonManagedKafkaEventSourceConfig,omitempty"`

	// The ARN of the Key Management Service (KMS) customer managed key that Lambda uses to encrypt your function's filter criteria.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximum_batching_window_in_seconds` expires or `batch_size` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`

	// - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	ParallelizationFactor int `json:"parallelizationFactor,omitempty" yaml:"parallelizationFactor,omitempty"`

	// For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include `self_managed_event_source`. Detailed below.
	SourceAccessConfigurations []types.Lambda_EventSourceMappingSourceAccessConfiguration `json:"sourceAccessConfigurations,omitempty" yaml:"sourceAccessConfigurations,omitempty"`

	// - (Optional) An Amazon SQS queue, Amazon SNS topic or Amazon S3 bucket (only available for Kafka sources) destination for failed records. Only available for stream sources (DynamoDB and Kinesis) and Kafka sources (Amazon MSK and Self-managed Apache Kafka). Detailed below.
	DestinationConfig types.Lambda_EventSourceMappingDestinationConfig `json:"destinationConfig,omitempty" yaml:"destinationConfig,omitempty"`

	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. The list must contain exactly one queue name.
	Queues string `json:"queues,omitempty" yaml:"queues,omitempty"`

	// Additional configuration block for Self Managed Kafka sources. Incompatible with "event_source_arn" and "amazon_managed_kafka_event_source_config". Detailed below.
	SelfManagedKafkaEventSourceConfig types.Lambda_EventSourceMappingSelfManagedKafkaEventSourceConfig `json:"selfManagedKafkaEventSourceConfig,omitempty" yaml:"selfManagedKafkaEventSourceConfig,omitempty"`

	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis, MQ and MSK, `10` for SQS.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// - (Optional) Configuration settings for a DocumentDB event source. Detailed below.
	DocumentDbEventSourceConfig types.Lambda_EventSourceMappingDocumentDbEventSourceConfig `json:"documentDbEventSourceConfig,omitempty" yaml:"documentDbEventSourceConfig,omitempty"`

	// The criteria to use for [event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
	FilterCriteria types.Lambda_EventSourceMappingFilterCriteria `json:"filterCriteria,omitempty" yaml:"filterCriteria,omitempty"`

	// A list of current response type enums applied to the event source mapping for [AWS Lambda checkpointing](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting). Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: `ReportBatchItemFailures`.
	FunctionResponseTypes []string `json:"functionResponseTypes,omitempty" yaml:"functionResponseTypes,omitempty"`

	// - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" yaml:"maximumRetryAttempts,omitempty"`
}
