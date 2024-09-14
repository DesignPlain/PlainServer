package types

type Certificateauthority_CertificateCertificateDescriptionCertFingerprint struct {
	/*
	   (Output)
	   The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate.
	*/
	Sha256Hash string `json:"sha256Hash,omitempty" yaml:"sha256Hash,omitempty"`
}
