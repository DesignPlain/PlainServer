package types

type Firebase_HostingCustomDomainRequiredDnsUpdateDesired struct {
	/*
	   Records on the domain
	   Structure is documented below.
	*/
	Records []Firebase_HostingCustomDomainRequiredDnsUpdateDesiredRecord `json:"records,omitempty" yaml:"records,omitempty"`

	// The domain name the record pertains to, e.g. `foo.bar.com.`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
