package types

type Firebase_HostingCustomDomainCertVerificationDnsDesired struct {
	// The domain name the record pertains to, e.g. `foo.bar.com.`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	/*
	   Records on the domain
	   Structure is documented below.
	*/
	Records []Firebase_HostingCustomDomainCertVerificationDnsDesiredRecord `json:"records,omitempty" yaml:"records,omitempty"`
}
