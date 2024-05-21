package types

type Secretmanager_SecretReplicationUserManaged struct {
	/*
	   The list of Replicas for this Secret. Cannot be empty.
	   Structure is documented below.
	*/
	Replicas []Secretmanager_SecretReplicationUserManagedReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}
