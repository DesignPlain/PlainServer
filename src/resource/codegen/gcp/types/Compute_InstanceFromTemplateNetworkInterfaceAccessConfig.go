package types

type Compute_InstanceFromTemplateNetworkInterfaceAccessConfig struct {
	// The IP address that is be 1:1 mapped to the instance's network ip.
	NatIp string `json:"natIp,omitempty" yaml:"natIp,omitempty"`

	// The networking tier used for configuring this instance. One of PREMIUM or STANDARD.
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`

	// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`
}
