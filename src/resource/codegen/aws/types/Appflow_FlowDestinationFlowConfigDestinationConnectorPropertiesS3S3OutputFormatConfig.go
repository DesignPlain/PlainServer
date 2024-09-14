package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig struct {
	// Aggregation settings that you can use to customize the output format of your flow data. See Aggregation Config for more details.
	AggregationConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig `json:"aggregationConfig,omitempty" yaml:"aggregationConfig,omitempty"`

	// File type that Amazon AppFlow places in the Amazon S3 bucket. Valid values are `CSV`, `JSON`, and `PARQUET`.
	FileType string `json:"fileType,omitempty" yaml:"fileType,omitempty"`

	// Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket. You can name folders according to the flow frequency and date. See Prefix Config for more details.
	PrefixConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig `json:"prefixConfig,omitempty" yaml:"prefixConfig,omitempty"`

	// Whether the data types from the source system need to be preserved (Only valid for `Parquet` file type)
	PreserveSourceDataTyping bool `json:"preserveSourceDataTyping,omitempty" yaml:"preserveSourceDataTyping,omitempty"`
}
