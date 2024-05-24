package types

type Index_ProviderIgnoreTags struct {
	// Resource tag key prefixes to ignore across all resources.
	KeyPrefixes []string `json:"keyPrefixes,omitempty" yaml:"keyPrefixes,omitempty"`

	// Resource tag keys to ignore across all resources.
	Keys []string `json:"keys,omitempty" yaml:"keys,omitempty"`
}
