package types

type Vmwareengine_ExternalAccessRuleDestinationIpRange struct {
	/*
	   The name of an `ExternalAddress` resource.

	   - - -
	*/
	ExternalAddress string `json:"externalAddress,omitempty" yaml:"externalAddress,omitempty"`

	// An IP address range in the CIDR format.
	IpAddressRange string `json:"ipAddressRange,omitempty" yaml:"ipAddressRange,omitempty"`
}
