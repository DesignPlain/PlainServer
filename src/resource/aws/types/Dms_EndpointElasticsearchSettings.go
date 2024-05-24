package types

type Dms_EndpointElasticsearchSettings struct {
	// Enable to migrate documentation using the documentation type `_doc`. OpenSearch and an Elasticsearch clusters only support the _doc documentation type in versions 7.x and later. The default value is `false`.
	UseNewMappingType bool `json:"useNewMappingType,omitempty" yaml:"useNewMappingType,omitempty"`

	// Endpoint for the OpenSearch cluster.
	EndpointUri string `json:"endpointUri,omitempty" yaml:"endpointUri,omitempty"`

	// Maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster. Default is `300`.
	ErrorRetryDuration int `json:"errorRetryDuration,omitempty" yaml:"errorRetryDuration,omitempty"`

	// Maximum percentage of records that can fail to be written before a full load operation stops. Default is `10`.
	FullLoadErrorPercentage int `json:"fullLoadErrorPercentage,omitempty" yaml:"fullLoadErrorPercentage,omitempty"`

	// ARN of the IAM Role with permissions to write to the OpenSearch cluster.
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`
}
