package types

type Firebase_HostingCustomDomainCertVerificationDns struct {
	/*
	   (Output)
	   The last time Hosting checked your CustomDomain's DNS records.
	*/
	CheckTime string `json:"checkTime,omitempty" yaml:"checkTime,omitempty"`

	// A text string to serve at the path.
	Desireds []Firebase_HostingCustomDomainCertVerificationDnsDesired `json:"desireds,omitempty" yaml:"desireds,omitempty"`

	/*
	   Whether Hosting was able to find the required file contents on the
	   specified path during its last check.
	*/
	Discovereds []Firebase_HostingCustomDomainCertVerificationDnsDiscovered `json:"discovereds,omitempty" yaml:"discovereds,omitempty"`
}
