package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration struct {
	// Specifies which serializer to use. You can choose either the ORC SerDe or the Parquet SerDe. See `serializer` block below for details.
	Serializer Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `json:"serializer,omitempty" yaml:"serializer,omitempty"`
}
