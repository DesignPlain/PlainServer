package types

type Signer_getSigningJobRevocationRecord struct {
	//
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	//
	RevokedAt string `json:"revokedAt,omitempty" yaml:"revokedAt,omitempty"`

	//
	RevokedBy string `json:"revokedBy,omitempty" yaml:"revokedBy,omitempty"`
}
