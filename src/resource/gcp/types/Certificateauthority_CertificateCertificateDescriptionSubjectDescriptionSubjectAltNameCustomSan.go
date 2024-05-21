package types

type Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSan struct {
	/*
	   Indicates whether or not this extension is critical (i.e., if the client does not know how to
	   handle this extension, the client should consider this to be an error).
	*/
	Critical bool `json:"critical,omitempty" yaml:"critical,omitempty"`

	/*
	   (Output)
	   Describes how some of the technical fields in a certificate should be populated.
	   Structure is documented below.
	*/
	ObectIds []Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSanObectId `json:"obectIds,omitempty" yaml:"obectIds,omitempty"`

	// The value of this X.509 extension. A base64-encoded string.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
