package ec2

import types "libds/aws/types"

type VpcPeeringConnectionAccepter struct {
	/*
	   A configuration block that describes [VPC Peering Connection]
	   (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	*/
	Accepter types.Ec2_VpcPeeringConnectionAccepterAccepter `json:"accepter,omitempty" yaml:"accepter,omitempty"`

	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept bool `json:"autoAccept,omitempty" yaml:"autoAccept,omitempty"`

	/*
	   A configuration block that describes [VPC Peering Connection]
	   (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	*/
	Requester types.Ec2_VpcPeeringConnectionAccepterRequester `json:"requester,omitempty" yaml:"requester,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`
}
