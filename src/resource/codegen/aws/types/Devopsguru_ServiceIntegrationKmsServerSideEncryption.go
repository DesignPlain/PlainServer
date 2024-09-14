package types

type Devopsguru_ServiceIntegrationKmsServerSideEncryption struct {
	// KMS key ID. This value can be a key ID, key ARN, alias name, or alias ARN.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies whether KMS integration is enabled. Valid values are `DISABLED` and `ENABLED`.
	OptInStatus string `json:"optInStatus,omitempty" yaml:"optInStatus,omitempty"`

	// Type of KMS key used. Valid values are `CUSTOMER_MANAGED_KEY` and `AWS_OWNED_KMS_KEY`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
