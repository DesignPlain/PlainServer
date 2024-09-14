package types

type Signer_SigningJobRevocationRecord struct {
	//
	RevokedAt string `json:"revokedAt,omitempty" yaml:"revokedAt,omitempty"`

	//
	RevokedBy string `json:"revokedBy,omitempty" yaml:"revokedBy,omitempty"`

	//
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
