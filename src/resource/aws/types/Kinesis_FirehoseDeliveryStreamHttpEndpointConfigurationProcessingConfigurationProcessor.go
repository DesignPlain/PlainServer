package types

type Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessor struct {
	// Specifies the processor parameters as multiple blocks. See `parameters` block below for details.
	Parameters []Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessorParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The type of processor. Valid Values: `RecordDeAggregation`, `Lambda`, `MetadataExtraction`, `AppendDelimiterToRecord`. Validation is done against [AWS SDK constants](https://docs.aws.amazon.com/sdk-for-go/api/service/firehose/#pkg-constants); so that values not explicitly listed may also work.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
