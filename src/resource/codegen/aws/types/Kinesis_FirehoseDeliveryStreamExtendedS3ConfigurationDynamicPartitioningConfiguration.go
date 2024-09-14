package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfiguration struct {
	// Enables or disables dynamic partitioning. Defaults to `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Total amount of seconds Firehose spends on retries. Valid values between 0 and 7200. Default is 300.

	   > --NOTE:-- You can enable dynamic partitioning only when you create a new delivery stream. Once you enable dynamic partitioning on a delivery stream, it cannot be disabled on this delivery stream. Therefore, the provider will recreate the resource whenever dynamic partitioning is enabled or disabled.
	*/
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`
}
