package types



type Secretmanager_getSecretsSecretReplicationUserManaged struct {
	/*
	   The list of Replicas for this Secret.
	   Structure is documented below.
	*/
	Replicas []Secretmanager_getSecretsSecretReplicationUserManagedReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}
