package types

type Opensearch_getDomainEncryptionAtRest struct {
	// Enabled disabled toggle for off-peak update window
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// KMS key id used to encrypt data at rest.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
