package types

type Certificateauthority_CertificateCertificateDescriptionSubjectKeyId struct {
	/*
	   (Output)
	   Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key.
	*/
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}
