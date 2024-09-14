package types

type Rds_InstanceMasterUserSecret struct {
	/*
	   The ARN for the KMS encryption key. If creating an
	   encrypted replica, set this to the destination KMS ARN.
	*/
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// The status of the secret. Valid Values: `creating` | `active` | `rotating` | `impaired`.
	SecretStatus string `json:"secretStatus,omitempty" yaml:"secretStatus,omitempty"`
}
