package types

type Ec2_getVpcIpamPoolCidrsIpamPoolCidr struct {
	// A network CIDR.
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// The provisioning state of that CIDR.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
