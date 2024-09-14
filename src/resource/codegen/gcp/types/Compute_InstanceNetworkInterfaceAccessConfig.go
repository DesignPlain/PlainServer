package types

type Compute_InstanceNetworkInterfaceAccessConfig struct {
	/*
	   The service-level to be provided for IPv6 traffic when the
	   subnet has an external subnet. Only PREMIUM or STANDARD tier is valid for IPv6.
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	/*
	   The domain name to be used when creating DNSv6
	   records for the external IPv6 ranges..
	*/
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`

	// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	/*
	   The IP address that will be 1:1 mapped to the instance's
	   network ip. If not given, one will be generated.
	*/
	NatIp string `json:"natIp,omitempty" yaml:"natIp,omitempty"`
}
