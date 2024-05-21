package types

type Dns_getKeysZoneSigningKeyDigest struct {
	// The base-16 encoded bytes of this digest. Suitable for use in a DS resource record.
	Digest string `json:"digest,omitempty" yaml:"digest,omitempty"`

	// Specifies the algorithm used to calculate this digest. Possible values are `sha1`, `sha256` and `sha384`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
