package types

type Ec2_getVpcPeeringConnectionPeerCidrBlockSet struct {
	// Primary CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`
}
