package types

type Compute_getInstanceNetworkInterfaceIpv6AccessConfig struct {
	// The DNS domain name for the public PTR record.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`

	// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	// The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
	ExternalIpv6 string `json:"externalIpv6,omitempty" yaml:"externalIpv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIpv6PrefixLength string `json:"externalIpv6PrefixLength,omitempty" yaml:"externalIpv6PrefixLength,omitempty"`

	// The name of the instance. One of `name` or `self_link` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The [networking tier][network-tier] used for configuring this instance. One of `PREMIUM` or `STANDARD`.
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`
}
