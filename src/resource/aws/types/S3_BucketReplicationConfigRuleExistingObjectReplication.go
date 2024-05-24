package types

type S3_BucketReplicationConfigRuleExistingObjectReplication struct {
	// Whether the existing objects should be replicated. Either `"Enabled"` or `"Disabled"`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
