package types

type Signer_SigningProfileRevocationRecord struct {
	//
	RevokedAt string `json:"revokedAt,omitempty" yaml:"revokedAt,omitempty"`

	//
	RevokedBy string `json:"revokedBy,omitempty" yaml:"revokedBy,omitempty"`

	//
	RevocationEffectiveFrom string `json:"revocationEffectiveFrom,omitempty" yaml:"revocationEffectiveFrom,omitempty"`
}
