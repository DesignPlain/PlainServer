package types

type Ecr_ReplicationConfigurationReplicationConfigurationRule struct {
	// the details of a replication destination. A maximum of 25 are allowed per `rule`. See Destination.
	Destinations []Ecr_ReplicationConfigurationReplicationConfigurationRuleDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// filters for a replication rule. See Repository Filter.
	RepositoryFilters []Ecr_ReplicationConfigurationReplicationConfigurationRuleRepositoryFilter `json:"repositoryFilters,omitempty" yaml:"repositoryFilters,omitempty"`
}
