package types

type Ecr_ReplicationConfigurationReplicationConfiguration struct {
	// The replication rules for a replication configuration. A maximum of 10 are allowed per `replication_configuration`. See Rule
	Rules []Ecr_ReplicationConfigurationReplicationConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
