package kinesis

import types "libds/aws/types"

type FirehoseDeliveryStream struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration options when `destination` is `elasticsearch`. See `elasticsearch_configuration` block below for details.
	ElasticsearchConfiguration types.Kinesis_FirehoseDeliveryStreamElasticsearchConfiguration `json:"elasticsearchConfiguration,omitempty" yaml:"elasticsearchConfiguration,omitempty"`

	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in. When using for WAF logging, name must be prefixed with `aws-waf-logs-`. See [AWS Documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-policies.html#waf-policies-logging-config) for more details.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration options when `destination` is `opensearchserverless`. See `opensearchserverless_configuration` block below for details.
	OpensearchserverlessConfiguration types.Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfiguration `json:"opensearchserverlessConfiguration,omitempty" yaml:"opensearchserverlessConfiguration,omitempty"`

	// Configuration options when `destination` is `redshift`. Requires the user to also specify an `s3_configuration` block. See `redshift_configuration` block below for details.
	RedshiftConfiguration types.Kinesis_FirehoseDeliveryStreamRedshiftConfiguration `json:"redshiftConfiguration,omitempty" yaml:"redshiftConfiguration,omitempty"`

	/*
	   Encrypt at rest options. See `server_side_encryption` block below for details.

	   --NOTE:-- Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
	*/
	ServerSideEncryption types.Kinesis_FirehoseDeliveryStreamServerSideEncryption `json:"serverSideEncryption,omitempty" yaml:"serverSideEncryption,omitempty"`

	// Configuration options when `destination` is `snowflake`. See `snowflake_configuration` block below for details.
	SnowflakeConfiguration types.Kinesis_FirehoseDeliveryStreamSnowflakeConfiguration `json:"snowflakeConfiguration,omitempty" yaml:"snowflakeConfiguration,omitempty"`

	// Configuration options when `destination` is `splunk`. See `splunk_configuration` block below for details.
	SplunkConfiguration types.Kinesis_FirehoseDeliveryStreamSplunkConfiguration `json:"splunkConfiguration,omitempty" yaml:"splunkConfiguration,omitempty"`

	// The Amazon Resource Name (ARN) specifying the Stream
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	//
	DestinationId string `json:"destinationId,omitempty" yaml:"destinationId,omitempty"`

	// Enhanced configuration options for the s3 destination. See `extended_s3_configuration` block below for details.
	ExtendedS3Configuration types.Kinesis_FirehoseDeliveryStreamExtendedS3Configuration `json:"extendedS3Configuration,omitempty" yaml:"extendedS3Configuration,omitempty"`

	//
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`

	// This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, `splunk`, `http_endpoint`, `opensearch`, `opensearchserverless` and `snowflake`.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The stream and role Amazon Resource Names (ARNs) for a Kinesis data stream used as the source for a delivery stream. See `kinesis_source_configuration` block below for details.
	KinesisSourceConfiguration types.Kinesis_FirehoseDeliveryStreamKinesisSourceConfiguration `json:"kinesisSourceConfiguration,omitempty" yaml:"kinesisSourceConfiguration,omitempty"`

	// The configuration for the Amazon MSK cluster to be used as the source for a delivery stream. See `msk_source_configuration` block below for details.
	MskSourceConfiguration types.Kinesis_FirehoseDeliveryStreamMskSourceConfiguration `json:"mskSourceConfiguration,omitempty" yaml:"mskSourceConfiguration,omitempty"`

	// Configuration options when `destination` is `opensearch`. See `opensearch_configuration` block below for details.
	OpensearchConfiguration types.Kinesis_FirehoseDeliveryStreamOpensearchConfiguration `json:"opensearchConfiguration,omitempty" yaml:"opensearchConfiguration,omitempty"`

	// Configuration options when `destination` is `http_endpoint`. Requires the user to also specify an `s3_configuration` block.  See `http_endpoint_configuration` block below for details.
	HttpEndpointConfiguration types.Kinesis_FirehoseDeliveryStreamHttpEndpointConfiguration `json:"httpEndpointConfiguration,omitempty" yaml:"httpEndpointConfiguration,omitempty"`
}
