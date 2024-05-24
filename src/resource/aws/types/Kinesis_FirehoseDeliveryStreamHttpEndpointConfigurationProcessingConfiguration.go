package types

type Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration struct {
	// Enables or disables data processing.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Specifies the data processors as multiple blocks. See `processors` block below for details.
	Processors []Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessor `json:"processors,omitempty" yaml:"processors,omitempty"`
}
