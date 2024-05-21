package types

type Certificateauthority_getAuthorityConfig struct {
	// Specifies some of the values in a certificate that are related to the subject.
	SubjectConfigs []Certificateauthority_getAuthorityConfigSubjectConfig `json:"subjectConfigs,omitempty" yaml:"subjectConfigs,omitempty"`

	// Describes how some of the technical X.509 fields in a certificate should be populated.
	X509Configs []Certificateauthority_getAuthorityConfigX509Config `json:"x509Configs,omitempty" yaml:"x509Configs,omitempty"`
}
