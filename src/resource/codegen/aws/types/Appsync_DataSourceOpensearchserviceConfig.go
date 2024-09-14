package types

type Appsync_DataSourceOpensearchserviceConfig struct {
	// HTTP endpoint of the OpenSearch domain.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// AWS region of the OpenSearch domain. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
