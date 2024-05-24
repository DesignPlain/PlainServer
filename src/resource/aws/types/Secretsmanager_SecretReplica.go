package types

type Secretsmanager_SecretReplica struct {
	// Region for replicating the secret.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Status can be `InProgress`, `Failed`, or `InSync`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`

	// ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Date that you last accessed the secret in the Region.
	LastAccessedDate string `json:"lastAccessedDate,omitempty" yaml:"lastAccessedDate,omitempty"`
}
