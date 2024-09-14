package types

type Rds_ClusterMasterUserSecret struct {
	// ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Amazon Resource Name (ARN) of the secret.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// Status of the secret. Valid Values: `creating` | `active` | `rotating` | `impaired`.
	SecretStatus string `json:"secretStatus,omitempty" yaml:"secretStatus,omitempty"`
}
