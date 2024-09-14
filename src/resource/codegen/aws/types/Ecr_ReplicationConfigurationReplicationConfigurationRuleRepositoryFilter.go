package types

type Ecr_ReplicationConfigurationReplicationConfigurationRuleRepositoryFilter struct {
	// The repository filter details.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// The repository filter type. The only supported value is `PREFIX_MATCH`, which is a repository name prefix specified with the filter parameter.
	FilterType string `json:"filterType,omitempty" yaml:"filterType,omitempty"`
}
