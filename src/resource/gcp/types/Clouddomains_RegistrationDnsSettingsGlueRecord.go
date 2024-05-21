package types

type Clouddomains_RegistrationDnsSettingsGlueRecord struct {
	// Required. Domain name of the host in Punycode format.
	HostName string `json:"hostName,omitempty" yaml:"hostName,omitempty"`

	/*
	   List of IPv4 addresses corresponding to this host in the standard decimal format (e.g. 198.51.100.1).
	   At least one of ipv4_address and ipv6_address must be set.
	*/
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty" yaml:"ipv4Addresses,omitempty"`

	/*
	   List of IPv4 addresses corresponding to this host in the standard decimal format (e.g. 198.51.100.1).
	   At least one of ipv4_address and ipv6_address must be set.
	*/
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses,omitempty"`
}
