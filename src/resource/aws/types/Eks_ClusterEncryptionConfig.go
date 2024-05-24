package types

type Eks_ClusterEncryptionConfig struct {
	// Configuration block with provider for encryption. Detailed below.
	Provider Eks_ClusterEncryptionConfigProvider `json:"provider,omitempty" yaml:"provider,omitempty"`

	// List of strings with resources to be encrypted. Valid values: `secrets`.
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`
}
