package types

type Ecr_ReplicationConfigurationReplicationConfigurationRuleDestination struct {
	// A Region to replicate to.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The account ID of the destination registry to replicate to.
	RegistryId string `json:"registryId,omitempty" yaml:"registryId,omitempty"`
}
