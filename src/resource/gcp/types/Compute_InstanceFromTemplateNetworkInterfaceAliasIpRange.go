package types

type Compute_InstanceFromTemplateNetworkInterfaceAliasIpRange struct {
	// The IP CIDR range represented by this alias IP range.
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	SubnetworkRangeName string `json:"subnetworkRangeName,omitempty" yaml:"subnetworkRangeName,omitempty"`
}
