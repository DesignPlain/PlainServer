package types

type Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationProcessingConfigurationProcessorParameter struct {
	// Parameter name. Valid Values: `LambdaArn`, `NumberOfRetries`, `MetadataExtractionQuery`, `JsonParsingEngine`, `RoleArn`, `BufferSizeInMBs`, `BufferIntervalInSeconds`, `SubRecordType`, `Delimiter`. Validation is done against [AWS SDK constants](https://docs.aws.amazon.com/sdk-for-go/api/service/firehose/#pkg-constants); so that values not explicitly listed may also work.
	ParameterName string `json:"parameterName,omitempty" yaml:"parameterName,omitempty"`

	/*
	   Parameter value. Must be between 1 and 512 length (inclusive). When providing a Lambda ARN, you should specify the resource version as well.

	   > --NOTE:-- Parameters with default values, including `NumberOfRetries`(default: 3), `RoleArn`(default: firehose role ARN), `BufferSizeInMBs`(default: 1), and `BufferIntervalInSeconds`(default: 60), are not stored in Pulumi state. To prevent perpetual differences, it is therefore recommended to only include parameters with non-default values.
	*/
	ParameterValue string `json:"parameterValue,omitempty" yaml:"parameterValue,omitempty"`
}
