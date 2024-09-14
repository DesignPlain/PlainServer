package types

type Firebase_HostingCustomDomainRequiredDnsUpdateDiscovered struct {
	// The domain name the record pertains to, e.g. `foo.bar.com.`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	/*
	   Records on the domain
	   Structure is documented below.
	*/
	Records []Firebase_HostingCustomDomainRequiredDnsUpdateDiscoveredRecord `json:"records,omitempty" yaml:"records,omitempty"`
}
