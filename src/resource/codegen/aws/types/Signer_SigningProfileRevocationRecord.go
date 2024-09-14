package types

type Signer_SigningProfileRevocationRecord struct {
	// The time when revocation becomes effective.
	RevocationEffectiveFrom string `json:"revocationEffectiveFrom,omitempty" yaml:"revocationEffectiveFrom,omitempty"`

	// The time when the signing profile was revoked.
	RevokedAt string `json:"revokedAt,omitempty" yaml:"revokedAt,omitempty"`

	// The identity of the revoker.
	RevokedBy string `json:"revokedBy,omitempty" yaml:"revokedBy,omitempty"`
}
