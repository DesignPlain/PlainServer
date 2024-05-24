package ecr

import types "DesignSphere_Server/src/resource/aws/types"

type ReplicationConfiguration struct {
	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration types.Ecr_ReplicationConfigurationReplicationConfiguration `json:"replicationConfiguration,omitempty" yaml:"replicationConfiguration,omitempty"`
}
