package types

type Certificateauthority_CertificateConfig struct {
	/*
	   A PublicKey describes a public key.
	   Structure is documented below.


	   <a name="nested_x509_config"></a>The `x509_config` block supports:
	*/
	PublicKey Certificateauthority_CertificateConfigPublicKey `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	/*
	   Specifies some of the values in a certificate that are related to the subject.
	   Structure is documented below.
	*/
	SubjectConfig Certificateauthority_CertificateConfigSubjectConfig `json:"subjectConfig,omitempty" yaml:"subjectConfig,omitempty"`

	/*
	   Describes how some of the technical X.509 fields in a certificate should be populated.
	   Structure is documented below.
	*/
	X509Config Certificateauthority_CertificateConfigX509Config `json:"x509Config,omitempty" yaml:"x509Config,omitempty"`
}
