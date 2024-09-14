package types

type Appsync_DataSourceElasticsearchConfig struct {
	// HTTP endpoint of the Elasticsearch domain.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// AWS region of Elasticsearch domain. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
