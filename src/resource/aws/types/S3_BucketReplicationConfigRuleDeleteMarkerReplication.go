package types

type S3_BucketReplicationConfigRuleDeleteMarkerReplication struct {
	// Whether delete markers should be replicated. Either `"Enabled"` or `"Disabled"`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
