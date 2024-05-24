package dynamodb

type TableReplica struct {
	// Storage class of the table replica. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`. If not used, the table replica will use the same class as the global table.
	TableClassOverride string `json:"tableClassOverride,omitempty" yaml:"tableClassOverride,omitempty"`

	// Map of tags to populate on the created table. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   ARN of the _main_ or global table which this resource will replicate.

	   Optional arguments:
	*/
	GlobalTableArn string `json:"globalTableArn,omitempty" yaml:"globalTableArn,omitempty"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. --Note:-- This attribute will _not_ be populated with the ARN of _default_ keys.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Whether to enable Point In Time Recovery for the replica. Default is `false`.
	PointInTimeRecovery bool `json:"pointInTimeRecovery,omitempty" yaml:"pointInTimeRecovery,omitempty"`
}
