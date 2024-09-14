package ecr

import types "libds/aws/types"

type ReplicationConfiguration struct {
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration types.Ecr_ReplicationConfigurationReplicationConfiguration `json:"replicationConfiguration,omitempty" yaml:"replicationConfiguration,omitempty"`
}
