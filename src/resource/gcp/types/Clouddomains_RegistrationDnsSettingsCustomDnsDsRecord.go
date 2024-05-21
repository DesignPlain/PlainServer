package types

type Clouddomains_RegistrationDnsSettingsCustomDnsDsRecord struct {
	// The algorithm used to generate the referenced DNSKEY.
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// The digest generated from the referenced DNSKEY.
	Digest string `json:"digest,omitempty" yaml:"digest,omitempty"`

	// The hash function used to generate the digest of the referenced DNSKEY.
	DigestType string `json:"digestType,omitempty" yaml:"digestType,omitempty"`

	// The key tag of the record. Must be set in range 0 -- 65535.
	KeyTag int `json:"keyTag,omitempty" yaml:"keyTag,omitempty"`
}
