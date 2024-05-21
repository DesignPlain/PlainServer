package types

type Certificateauthority_CertificateConfigSubjectConfig struct {
	/*
	   Contains distinguished name fields such as the location and organization.
	   Structure is documented below.
	*/
	Subject Certificateauthority_CertificateConfigSubjectConfigSubject `json:"subject,omitempty" yaml:"subject,omitempty"`

	/*
	   The subject alternative name fields.
	   Structure is documented below.
	*/
	SubjectAltName Certificateauthority_CertificateConfigSubjectConfigSubjectAltName `json:"subjectAltName,omitempty" yaml:"subjectAltName,omitempty"`
}
