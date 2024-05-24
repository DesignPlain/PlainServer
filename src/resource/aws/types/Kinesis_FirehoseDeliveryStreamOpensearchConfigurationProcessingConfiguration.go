package types

type Kinesis_FirehoseDeliveryStreamOpensearchConfigurationProcessingConfiguration struct {
	// Specifies the data processors as multiple blocks. See `processors` block below for details.
	Processors []Kinesis_FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessor `json:"processors,omitempty" yaml:"processors,omitempty"`

	// Enables or disables data processing.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
