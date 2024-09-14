package types

type Certificateauthority_CertificateRevocationDetail struct {
	/*
	   (Output)
	   Indicates why a Certificate was revoked.
	*/
	RevocationState string `json:"revocationState,omitempty" yaml:"revocationState,omitempty"`

	/*
	   (Output)
	   The time at which this Certificate was revoked.
	*/
	RevocationTime string `json:"revocationTime,omitempty" yaml:"revocationTime,omitempty"`
}
