package types

type S3_BucketReplicationConfigRuleDestinationReplicationTime struct {
	// Configuration block specifying the time by which replication should be complete for all objects and operations on objects. See below.
	Time S3_BucketReplicationConfigRuleDestinationReplicationTimeTime `json:"time,omitempty" yaml:"time,omitempty"`

	// Status of the Replication Time Control. Either `"Enabled"` or `"Disabled"`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
