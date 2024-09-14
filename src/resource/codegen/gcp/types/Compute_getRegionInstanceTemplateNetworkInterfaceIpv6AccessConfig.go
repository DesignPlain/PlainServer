package types

type Compute_getRegionInstanceTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The name of the instance template. One of `name` or `filter` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The [networking tier][network-tier] used for configuring
	   this instance template. This field can take the following values: PREMIUM or
	   STANDARD. If this field is not specified, it is assumed to be PREMIUM.
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`

	// The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.
	ExternalIpv6 string `json:"externalIpv6,omitempty" yaml:"externalIpv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIpv6PrefixLength string `json:"externalIpv6PrefixLength,omitempty" yaml:"externalIpv6PrefixLength,omitempty"`
}
