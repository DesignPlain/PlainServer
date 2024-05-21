package types

type Certificateauthority_AuthorityConfig struct {
	/*
	   Specifies some of the values in a certificate that are related to the subject.
	   Structure is documented below.


	   <a name="nested_x509_config"></a>The `x509_config` block supports:
	*/
	SubjectConfig Certificateauthority_AuthorityConfigSubjectConfig `json:"subjectConfig,omitempty" yaml:"subjectConfig,omitempty"`

	/*
	   Describes how some of the technical X.509 fields in a certificate should be populated.
	   Structure is documented below.
	*/
	X509Config Certificateauthority_AuthorityConfigX509Config `json:"x509Config,omitempty" yaml:"x509Config,omitempty"`
}
