package types

type S3_BucketReplicationConfigurationRule struct {
	// Specifies special object selection criteria (documented below).
	SourceSelectionCriteria S3_BucketReplicationConfigurationRuleSourceSelectionCriteria `json:"sourceSelectionCriteria,omitempty" yaml:"sourceSelectionCriteria,omitempty"`

	/*
	   The status of the rule. Either `Enabled` or `Disabled`. The rule is ignored if status is not Enabled.

	   > --NOTE:-- Replication to multiple destination buckets requires that `priority` is specified in the `rules` object. If the corresponding rule requires no filter, an empty configuration block `filter {}` must be specified.
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Whether delete markers are replicated. The only valid value is `Enabled`. To disable, omit this argument. This argument is only valid with V2 replication configurations (i.e., when `filter` is used).
	DeleteMarkerReplicationStatus string `json:"deleteMarkerReplicationStatus,omitempty" yaml:"deleteMarkerReplicationStatus,omitempty"`

	// Specifies the destination for the rule (documented below).
	Destination S3_BucketReplicationConfigurationRuleDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies (documented below).
	Filter S3_BucketReplicationConfigurationRuleFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Object keyname prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The priority associated with the rule. Priority should only be set if `filter` is configured. If not provided, defaults to `0`. Priority must be unique between multiple rules.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
