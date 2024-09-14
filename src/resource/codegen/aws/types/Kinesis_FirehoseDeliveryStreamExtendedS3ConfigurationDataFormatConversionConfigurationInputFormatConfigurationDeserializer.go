package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	// Specifies the native Hive / HCatalog JsonSerDe. More details below. See `hive_json_ser_de` block below for details.
	HiveJsonSerDe Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe `json:"hiveJsonSerDe,omitempty" yaml:"hiveJsonSerDe,omitempty"`

	// Specifies the OpenX SerDe. See `open_x_json_ser_de` block below for details.
	OpenXJsonSerDe Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe `json:"openXJsonSerDe,omitempty" yaml:"openXJsonSerDe,omitempty"`
}
