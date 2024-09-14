package types

type Dynamodb_TableReplica struct {
	// Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
	StreamLabel string `json:"streamLabel,omitempty" yaml:"streamLabel,omitempty"`

	// ARN of the table
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. --Note:-- This attribute will _not_ be populated with the ARN of _default_ keys.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Whether to enable Point In Time Recovery for the replica. Default is `false`.
	PointInTimeRecovery bool `json:"pointInTimeRecovery,omitempty" yaml:"pointInTimeRecovery,omitempty"`

	// Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
	PropagateTags bool `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`

	// Region name of the replica.
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`

	// ARN of the Table Stream. Only available when `stream_enabled = true`
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`
}
