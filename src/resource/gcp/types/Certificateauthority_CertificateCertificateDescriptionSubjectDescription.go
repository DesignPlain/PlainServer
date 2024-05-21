package types

type Certificateauthority_CertificateCertificateDescriptionSubjectDescription struct {
	/*
	   (Output)
	   The serial number encoded in lowercase hexadecimal.
	*/
	HexSerialNumber string `json:"hexSerialNumber,omitempty" yaml:"hexSerialNumber,omitempty"`

	/*
	   The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	   "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	   fractional digits, terminated by 's'. Example: "3.5s".
	*/
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`

	/*
	   (Output)
	   The time at which the certificate expires.
	*/
	NotAfterTime string `json:"notAfterTime,omitempty" yaml:"notAfterTime,omitempty"`

	/*
	   (Output)
	   The time at which the certificate becomes valid.
	*/
	NotBeforeTime string `json:"notBeforeTime,omitempty" yaml:"notBeforeTime,omitempty"`

	/*
	   The subject alternative name fields.
	   Structure is documented below.
	*/
	SubjectAltNames []Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubjectAltName `json:"subjectAltNames,omitempty" yaml:"subjectAltNames,omitempty"`

	/*
	   Contains distinguished name fields such as the location and organization.
	   Structure is documented below.
	*/
	Subjects []Certificateauthority_CertificateCertificateDescriptionSubjectDescriptionSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
}
