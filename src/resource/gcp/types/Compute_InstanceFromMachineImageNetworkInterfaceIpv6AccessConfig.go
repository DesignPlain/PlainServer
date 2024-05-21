package types

type Compute_InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`

	// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	// The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
	ExternalIpv6 string `json:"externalIpv6,omitempty" yaml:"externalIpv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIpv6PrefixLength string `json:"externalIpv6PrefixLength,omitempty" yaml:"externalIpv6PrefixLength,omitempty"`

	/*
	   A unique name for the resource, required by GCE.
	   Changing this forces a new resource to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
