package types

type Appsync_DataSourceDynamodbConfig struct {
	// The DeltaSyncConfig for a versioned data source. See Delta Sync Config
	DeltaSyncConfig Appsync_DataSourceDynamodbConfigDeltaSyncConfig `json:"deltaSyncConfig,omitempty" yaml:"deltaSyncConfig,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Name of the DynamoDB table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// Set to `true` to use Amazon Cognito credentials with this data source.
	UseCallerCredentials bool `json:"useCallerCredentials,omitempty" yaml:"useCallerCredentials,omitempty"`

	// Detects Conflict Detection and Resolution with this data source.
	Versioned bool `json:"versioned,omitempty" yaml:"versioned,omitempty"`
}
