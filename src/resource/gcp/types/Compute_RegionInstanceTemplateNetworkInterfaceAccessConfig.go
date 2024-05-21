package types

type Compute_RegionInstanceTemplateNetworkInterfaceAccessConfig struct {
	/*
	   The IP address that will be 1:1 mapped to the instance's
	   network ip. If not given, one will be generated.
	*/
	NatIp string `json:"natIp,omitempty" yaml:"natIp,omitempty"`

	/*
	   The service-level to be provided for IPv6 traffic when the
	   subnet has an external subnet. Only PREMIUM and STANDARD tier is valid for IPv6.
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record.The DNS domain name for the public PTR record.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`
}
