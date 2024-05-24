package types

type Appsync_DataSourceRelationalDatabaseConfigHttpEndpointConfig struct {
	// AWS secret store ARN for database credentials.
	AwsSecretStoreArn string `json:"awsSecretStoreArn,omitempty" yaml:"awsSecretStoreArn,omitempty"`

	// Logical database name.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Amazon RDS cluster identifier.
	DbClusterIdentifier string `json:"dbClusterIdentifier,omitempty" yaml:"dbClusterIdentifier,omitempty"`

	// AWS Region for RDS HTTP endpoint. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Logical schema name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
