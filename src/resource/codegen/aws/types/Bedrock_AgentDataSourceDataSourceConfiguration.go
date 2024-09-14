package types

type Bedrock_AgentDataSourceDataSourceConfiguration struct {
	// Type of storage for the data source. Valid values: `S3`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Details about the configuration of the S3 object containing the data source. See `s3_data_source_configuration` block for details.
	S3Configuration Bedrock_AgentDataSourceDataSourceConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`
}
