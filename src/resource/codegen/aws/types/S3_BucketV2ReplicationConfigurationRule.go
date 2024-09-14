package types

type S3_BucketV2ReplicationConfigurationRule struct {
	// Whether delete markers are replicated. The only valid value is `Enabled`. To disable, omit this argument. This argument is only valid with V2 replication configurations (i.e., when `filter` is used).
	DeleteMarkerReplicationStatus string `json:"deleteMarkerReplicationStatus,omitempty" yaml:"deleteMarkerReplicationStatus,omitempty"`

	// Specifies the destination for the rule (documented below).
	Destinations []S3_BucketV2ReplicationConfigurationRuleDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies (documented below).
	Filters []S3_BucketV2ReplicationConfigurationRuleFilter `json:"filters,omitempty" yaml:"filters,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Object keyname prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Priority associated with the rule. Priority should only be set if `filter` is configured. If not provided, defaults to `0`. Priority must be unique between multiple rules.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Specifies special object selection criteria (documented below).
	SourceSelectionCriterias []S3_BucketV2ReplicationConfigurationRuleSourceSelectionCriteria `json:"sourceSelectionCriterias,omitempty" yaml:"sourceSelectionCriterias,omitempty"`

	// Status of the rule. Either `Enabled` or `Disabled`. The rule is ignored if status is not Enabled.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
