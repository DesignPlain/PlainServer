package lightsail

type LbCertificate struct {
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName string `json:"lbName,omitempty" yaml:"lbName,omitempty"`

	// The SSL/TLS certificate name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`
}
