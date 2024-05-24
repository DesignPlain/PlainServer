package types

type Customerprofiles_DomainRuleBasedMatchingConflictResolution struct {
	// How the auto-merging process should resolve conflicts between different profiles. Valid values are `RECENCY` and `SOURCE`
	ConflictResolvingModel string `json:"conflictResolvingModel,omitempty" yaml:"conflictResolvingModel,omitempty"`

	// The `ObjectType` name that is used to resolve profile merging conflicts when choosing `SOURCE` as the `ConflictResolvingModel`.
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`
}
