package types

type Secretmanager_getSecretReplicationUserManaged struct {
	// The list of Replicas for this Secret. Cannot be empty.
	Replicas []Secretmanager_getSecretReplicationUserManagedReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}
