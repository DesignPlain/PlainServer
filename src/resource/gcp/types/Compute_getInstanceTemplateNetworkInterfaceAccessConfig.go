package types

type Compute_getInstanceTemplateNetworkInterfaceAccessConfig struct {
	/*
	   The IP address that will be 1:1 mapped to the instance's
	   network ip. If not given, one will be generated.
	*/
	NatIp string `json:"natIp,omitempty" yaml:"natIp,omitempty"`

	/*
	   The [networking tier][network-tier] used for configuring
	   this instance template. This field can take the following values: PREMIUM or
	   STANDARD. If this field is not specified, it is assumed to be PREMIUM.
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	// The DNS domain name for the public PTR record.The DNS domain name for the public PTR record.
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" yaml:"publicPtrDomainName,omitempty"`
}
