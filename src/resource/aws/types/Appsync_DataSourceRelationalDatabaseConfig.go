package types

type Appsync_DataSourceRelationalDatabaseConfig struct {
	// Amazon RDS HTTP endpoint configuration. See HTTP Endpoint Config.
	HttpEndpointConfig Appsync_DataSourceRelationalDatabaseConfigHttpEndpointConfig `json:"httpEndpointConfig,omitempty" yaml:"httpEndpointConfig,omitempty"`

	// Source type for the relational database. Valid values: `RDS_HTTP_ENDPOINT`.
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`
}
