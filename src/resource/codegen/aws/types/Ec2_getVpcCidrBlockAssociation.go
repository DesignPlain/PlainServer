package types

type Ec2_getVpcCidrBlockAssociation struct {
	// Association ID for the IPv4 CIDR block.
	AssociationId string `json:"associationId,omitempty" yaml:"associationId,omitempty"`

	// Cidr block of the desired VPC.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	/*
	   Current state of the desired VPC.
	   Can be either `"pending"` or `"available"`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
