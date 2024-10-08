package types

type Dns_getKeysKeySigningKey struct {
	// Base64 encoded public half of this key.
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	// String mnemonic specifying the DNSSEC algorithm of this key. Immutable after creation time. Possible values are `ecdsap256sha256`, `ecdsap384sha384`, `rsasha1`, `rsasha256`, and `rsasha512`.
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// The time that this resource was created in the control plane. This is in RFC3339 text format.
	CreationTime string `json:"creationTime,omitempty" yaml:"creationTime,omitempty"`

	// A mutable string of at most 1024 characters associated with this resource for the user's convenience.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The DS record based on the KSK record. This is used when [delegating](https://cloud.google.com/dns/docs/dnssec-advanced#subdelegation) DNSSEC-signed subdomains.
	DsRecord string `json:"dsRecord,omitempty" yaml:"dsRecord,omitempty"`

	// Length of the key in bits. Specified at creation time then immutable.
	KeyLength int `json:"keyLength,omitempty" yaml:"keyLength,omitempty"`

	// The key tag is a non-cryptographic hash of the a DNSKEY resource record associated with this DnsKey. The key tag can be used to identify a DNSKEY more quickly (but it is not a unique identifier). In particular, the key tag is used in a parent zone's DS record to point at the DNSKEY in this child ManagedZone. The key tag is a number in the range [0, 65535] and the algorithm to calculate it is specified in RFC4034 Appendix B.
	KeyTag int `json:"keyTag,omitempty" yaml:"keyTag,omitempty"`

	// A list of cryptographic hashes of the DNSKEY resource record associated with this DnsKey. These digests are needed to construct a DS record that points at this DNS key. Each contains:
	Digests []Dns_getKeysKeySigningKeyDigest `json:"digests,omitempty" yaml:"digests,omitempty"`

	// Unique identifier for the resource; defined by the server.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Active keys will be used to sign subsequent changes to the ManagedZone. Inactive keys will still be present as DNSKEY Resource Records for the use of resolvers validating existing signatures.
	IsActive bool `json:"isActive,omitempty" yaml:"isActive,omitempty"`
}
