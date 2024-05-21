package types

type Clouddomains_RegistrationDnsSettingsCustomDns struct {
	/*
	   Required. A list of name servers that store the DNS zone for this domain. Each name server is a domain
	   name, with Unicode domain names expressed in Punycode format.
	*/
	NameServers []string `json:"nameServers,omitempty" yaml:"nameServers,omitempty"`

	/*
	   The list of DS records for this domain, which are used to enable DNSSEC. The domain's DNS provider can provide
	   the values to set here. If this field is empty, DNSSEC is disabled.
	   Structure is documented below.
	*/
	DsRecords []Clouddomains_RegistrationDnsSettingsCustomDnsDsRecord `json:"dsRecords,omitempty" yaml:"dsRecords,omitempty"`
}
