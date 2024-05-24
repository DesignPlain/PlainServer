package types

type S3_BucketReplicationConfigRuleDestinationMetrics struct {
	// Configuration block that specifies the time threshold for emitting the `s3:Replication:OperationMissedThreshold` event. See below.
	EventThreshold S3_BucketReplicationConfigRuleDestinationMetricsEventThreshold `json:"eventThreshold,omitempty" yaml:"eventThreshold,omitempty"`

	// Status of the Destination Metrics. Either `"Enabled"` or `"Disabled"`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
