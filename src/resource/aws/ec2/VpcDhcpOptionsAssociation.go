package ec2

type VpcDhcpOptionsAssociation struct {
	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId string `json:"dhcpOptionsId,omitempty" yaml:"dhcpOptionsId,omitempty"`

	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
