package types

type Acm_CertificateValidationOption struct {
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Domain name that you want ACM to use to send you validation emails. This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the `domain_name` value or a superdomain of the `domain_name` value. For example, if you request a certificate for `"testing.example.com"`, you can specify `"example.com"` for this value.
	ValidationDomain string `json:"validationDomain,omitempty" yaml:"validationDomain,omitempty"`
}
