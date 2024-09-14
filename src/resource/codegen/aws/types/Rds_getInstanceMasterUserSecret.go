package types

type Rds_getInstanceMasterUserSecret struct {
	// The Amazon Web Services KMS key identifier that is used to encrypt the secret.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// The status of the secret. Valid Values: `creating` | `active` | `rotating` | `impaired`.
	SecretStatus string `json:"secretStatus,omitempty" yaml:"secretStatus,omitempty"`
}
