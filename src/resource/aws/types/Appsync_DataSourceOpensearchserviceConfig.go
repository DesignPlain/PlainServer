package types

type Appsync_DataSourceOpensearchserviceConfig struct {
	// HTTP endpoint of the Elasticsearch domain.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
