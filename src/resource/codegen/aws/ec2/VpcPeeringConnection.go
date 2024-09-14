package ec2

import types "libds/aws/types"

type VpcPeeringConnection struct {
	// Accept the peering (both VPCs need to be in the same AWS account and region).
	AutoAccept bool `json:"autoAccept,omitempty" yaml:"autoAccept,omitempty"`

	/*
	   The AWS account ID of the target peer VPC.
	   Defaults to the account ID the [AWS provider][1] is currently connected to, so must be managed if connecting cross-account.
	*/
	PeerOwnerId string `json:"peerOwnerId,omitempty" yaml:"peerOwnerId,omitempty"`

	/*
	   The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
	   and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	*/
	PeerRegion string `json:"peerRegion,omitempty" yaml:"peerRegion,omitempty"`

	// The ID of the target VPC with which you are creating the VPC Peering Connection.
	PeerVpcId string `json:"peerVpcId,omitempty" yaml:"peerVpcId,omitempty"`

	/*
	   A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
	   the peering connection (a maximum of one).
	*/
	Requester types.Ec2_VpcPeeringConnectionRequester `json:"requester,omitempty" yaml:"requester,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the requester VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	/*
	   An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
	   the peering connection (a maximum of one).
	*/
	Accepter types.Ec2_VpcPeeringConnectionAccepter `json:"accepter,omitempty" yaml:"accepter,omitempty"`
}
