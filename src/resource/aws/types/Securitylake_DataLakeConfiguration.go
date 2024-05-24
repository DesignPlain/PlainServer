package types

type Securitylake_DataLakeConfiguration struct {
	// Provides encryption details of Amazon Security Lake object.
	EncryptionConfigurations []Securitylake_DataLakeConfigurationEncryptionConfiguration `json:"encryptionConfigurations,omitempty" yaml:"encryptionConfigurations,omitempty"`

	// Provides lifecycle details of Amazon Security Lake object.
	LifecycleConfiguration Securitylake_DataLakeConfigurationLifecycleConfiguration `json:"lifecycleConfiguration,omitempty" yaml:"lifecycleConfiguration,omitempty"`

	// The AWS Regions where Security Lake is automatically enabled.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Provides replication details of Amazon Security Lake object.
	ReplicationConfiguration Securitylake_DataLakeConfigurationReplicationConfiguration `json:"replicationConfiguration,omitempty" yaml:"replicationConfiguration,omitempty"`
}
