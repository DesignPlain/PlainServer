package types

type Compute_getSubnetworkSecondaryIpRange struct {
	/*
	   The range of IP addresses belonging to this subnetwork
	   secondary range.
	*/
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	/*
	   The name associated with this subnetwork secondary range, used
	   when adding an alias IP range to a VM instance.
	*/
	RangeName string `json:"rangeName,omitempty" yaml:"rangeName,omitempty"`
}
