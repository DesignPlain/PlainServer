package types

type S3_BucketReplicationConfigRule struct {
	// Priority associated with the rule. Priority should only be set if `filter` is configured. If not provided, defaults to `0`. Priority must be unique between multiple rules.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Status of the rule. Either `"Enabled"` or `"Disabled"`. The rule is ignored if status is not "Enabled".
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Specifies special object selection criteria. See below.
	SourceSelectionCriteria S3_BucketReplicationConfigRuleSourceSelectionCriteria `json:"sourceSelectionCriteria,omitempty" yaml:"sourceSelectionCriteria,omitempty"`

	// Whether delete markers are replicated. This argument is only valid with V2 replication configurations (i.e., when `filter` is used)documented below.
	DeleteMarkerReplication S3_BucketReplicationConfigRuleDeleteMarkerReplication `json:"deleteMarkerReplication,omitempty" yaml:"deleteMarkerReplication,omitempty"`

	// Specifies the destination for the rule. See below.
	Destination S3_BucketReplicationConfigRuleDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Replicate existing objects in the source bucket according to the rule configurations. See below.
	ExistingObjectReplication S3_BucketReplicationConfigRuleExistingObjectReplication `json:"existingObjectReplication,omitempty" yaml:"existingObjectReplication,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies. See below. If not specified, the `rule` will default to using `prefix`.
	Filter S3_BucketReplicationConfigRuleFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string (`""`) if `filter` is not specified.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
