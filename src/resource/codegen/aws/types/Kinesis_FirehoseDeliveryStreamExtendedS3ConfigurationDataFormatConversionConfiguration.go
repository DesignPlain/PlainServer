package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfiguration struct {
	// Specifies the serializer that you want Kinesis Data Firehose to use to convert the format of your data to the Parquet or ORC format. See `output_format_configuration` block below for details.
	OutputFormatConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration `json:"outputFormatConfiguration,omitempty" yaml:"outputFormatConfiguration,omitempty"`

	// Specifies the AWS Glue Data Catalog table that contains the column information. See `schema_configuration` block below for details.
	SchemaConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration `json:"schemaConfiguration,omitempty" yaml:"schemaConfiguration,omitempty"`

	// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Specifies the deserializer that you want Kinesis Data Firehose to use to convert the format of your data from JSON. See `input_format_configuration` block below for details.
	InputFormatConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration `json:"inputFormatConfiguration,omitempty" yaml:"inputFormatConfiguration,omitempty"`
}
