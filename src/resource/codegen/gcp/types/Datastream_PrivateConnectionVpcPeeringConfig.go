package types

type Datastream_PrivateConnectionVpcPeeringConfig struct {
	/*
	   A free subnet for peering. (CIDR of /29)

	   - - -
	*/
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	/*
	   Fully qualified name of the VPC that Datastream will peer to.
	   Format: projects/{project}/global/{networks}/{name}
	*/
	Vpc string `json:"vpc,omitempty" yaml:"vpc,omitempty"`
}
