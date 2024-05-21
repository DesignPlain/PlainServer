package types

type Certificateauthority_CertificateCertificateDescription struct {
	/*
	   (Output)
	   A structured description of the issued X.509 certificate.
	   Structure is documented below.
	*/
	X509Descriptions []Certificateauthority_CertificateCertificateDescriptionX509Description `json:"x509Descriptions,omitempty" yaml:"x509Descriptions,omitempty"`

	/*
	   (Output)
	   Describes lists of issuer CA certificate URLs that appear in the "Authority Information Access" extension in the certificate.
	*/
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty" yaml:"aiaIssuingCertificateUrls,omitempty"`

	/*
	   (Output)
	   Identifies the subjectKeyId of the parent certificate, per https://tools.ietf.org/html/rfc5280#section-4.2.1.1
	   Structure is documented below.
	*/
	AuthorityKeyIds []Certificateauthority_CertificateCertificateDescriptionAuthorityKeyId `json:"authorityKeyIds,omitempty" yaml:"authorityKeyIds,omitempty"`

	/*
	   (Output)
	   The hash of the x.509 certificate.
	   Structure is documented below.
	*/
	CertFingerprints []Certificateauthority_CertificateCertificateDescriptionCertFingerprint `json:"certFingerprints,omitempty" yaml:"certFingerprints,omitempty"`

	/*
	   (Output)
	   Describes a list of locations to obtain CRL information, i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13
	*/
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty" yaml:"crlDistributionPoints,omitempty"`

	/*
	   A PublicKey describes a public key.
	   Structure is documented below.


	   <a name="nested_x509_config"></a>The `x509_config` block supports:
	*/
	PublicKeys []Certificateauthority_CertificateCertificateDescriptionPublicKey `json:"publicKeys,omitempty" yaml:"publicKeys,omitempty"`

	/*
	   (Output)
	   Describes some of the values in a certificate that are related to the subject and lifetime.
	   Structure is documented below.
	*/
	SubjectDescriptions []Certificateauthority_CertificateCertificateDescriptionSubjectDescription `json:"subjectDescriptions,omitempty" yaml:"subjectDescriptions,omitempty"`

	/*
	   (Output)
	   Provides a means of identifiying certificates that contain a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2.
	   Structure is documented below.
	*/
	SubjectKeyIds []Certificateauthority_CertificateCertificateDescriptionSubjectKeyId `json:"subjectKeyIds,omitempty" yaml:"subjectKeyIds,omitempty"`
}
