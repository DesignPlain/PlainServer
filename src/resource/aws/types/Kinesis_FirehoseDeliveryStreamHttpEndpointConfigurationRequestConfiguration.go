package types

type Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfiguration struct {
	// Describes the metadata sent to the HTTP endpoint destination. See `common_attributes` block below for details.
	CommonAttributes []Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttribute `json:"commonAttributes,omitempty" yaml:"commonAttributes,omitempty"`

	// Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination. Valid values are `NONE` and `GZIP`.  Default value is `NONE`.
	ContentEncoding string `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`
}
