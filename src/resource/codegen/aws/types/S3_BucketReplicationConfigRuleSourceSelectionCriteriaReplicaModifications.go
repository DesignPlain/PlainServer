package types

type S3_BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications struct {
	// Whether the existing objects should be replicated. Either `"Enabled"` or `"Disabled"`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
