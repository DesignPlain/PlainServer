package types

type S3_BucketReplicationConfigurationRuleDestinationMetrics struct {
	// Threshold within which objects are to be replicated. The only valid value is `15`.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`

	// The status of replication metrics. Either `Enabled` or `Disabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
