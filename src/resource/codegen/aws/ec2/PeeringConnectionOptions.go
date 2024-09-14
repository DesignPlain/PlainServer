package ec2

import types "libds/aws/types"

type PeeringConnectionOptions struct {
	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter types.Ec2_PeeringConnectionOptionsAccepter `json:"accepter,omitempty" yaml:"accepter,omitempty"`

	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester types.Ec2_PeeringConnectionOptionsRequester `json:"requester,omitempty" yaml:"requester,omitempty"`

	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`
}
