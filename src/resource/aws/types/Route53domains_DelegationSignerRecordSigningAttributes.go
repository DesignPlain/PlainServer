package types

type Route53domains_DelegationSignerRecordSigningAttributes struct {
	// Algorithm which was used to generate the digest from the public key.
	Algorithm int `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// Defines the type of key. It can be either a KSK (key-signing-key, value `257`) or ZSK (zone-signing-key, value `256`).
	Flags int `json:"flags,omitempty" yaml:"flags,omitempty"`

	// The base64-encoded public key part of the key pair that is passed to the registry.
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
}
