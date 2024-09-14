package types

type Databasemigrationservice_PrivateConnectionVpcPeeringConfig struct {
	/*
	   A free subnet for peering. (CIDR of /29)

	   - - -
	*/
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	/*
	   Fully qualified name of the VPC that Database Migration Service will peer to.
	   Format: projects/{project}/global/{networks}/{name}
	*/
	VpcName string `json:"vpcName,omitempty" yaml:"vpcName,omitempty"`
}
