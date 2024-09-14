package types

type Index_ProviderIgnoreTags struct {
	// Resource tag keys to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEYS environment variable.
	Keys []string `json:"keys,omitempty" yaml:"keys,omitempty"`

	// Resource tag key prefixes to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEY_PREFIXES environment variable.
	KeyPrefixes []string `json:"keyPrefixes,omitempty" yaml:"keyPrefixes,omitempty"`
}
