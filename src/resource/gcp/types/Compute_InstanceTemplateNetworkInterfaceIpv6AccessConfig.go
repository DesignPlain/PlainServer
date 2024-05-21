package types

type Compute_InstanceTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.
	ExternalIpv6 string `json:"externalIpv6,omitempty" yaml:"externalIpv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIpv6PrefixLength string `json:"externalIpv6PrefixLength,omitempty" yaml:"externalIpv6PrefixLength,omitempty"`

	/*
	   The name of the instance template. If you leave
	   this blank, the provider will auto-generate a unique name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The service-level to be provided for IPv6 traffic when the
	   subnet has an external subnet. Only PREMIUM and STANDARD tier is valid for IPv6.
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`
}
