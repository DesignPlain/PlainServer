package types

type Verifiedaccess_InstanceLoggingConfigurationAccessLogsKinesisDataFirehose struct {
	// The name of the delivery stream.
	DeliveryStream string `json:"deliveryStream,omitempty" yaml:"deliveryStream,omitempty"`

	// Indicates whether logging is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
