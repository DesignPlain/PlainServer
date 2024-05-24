package types

type Signer_getSigningProfileRevocationRecord struct {
	//
	RevocationEffectiveFrom string `json:"revocationEffectiveFrom,omitempty" yaml:"revocationEffectiveFrom,omitempty"`

	//
	RevokedAt string `json:"revokedAt,omitempty" yaml:"revokedAt,omitempty"`

	//
	RevokedBy string `json:"revokedBy,omitempty" yaml:"revokedBy,omitempty"`
}
