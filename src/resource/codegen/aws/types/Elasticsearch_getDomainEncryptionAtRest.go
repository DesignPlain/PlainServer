package types

type Elasticsearch_getDomainEncryptionAtRest struct {
	// Whether node to node encryption is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The KMS key id used to encrypt data at rest.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
