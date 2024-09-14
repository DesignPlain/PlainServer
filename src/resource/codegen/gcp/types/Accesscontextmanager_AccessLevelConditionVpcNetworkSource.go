package types

type Accesscontextmanager_AccessLevelConditionVpcNetworkSource struct {
	/*
	   Sub networks within a VPC network.
	   Structure is documented below.
	*/
	VpcSubnetwork Accesscontextmanager_AccessLevelConditionVpcNetworkSourceVpcSubnetwork `json:"vpcSubnetwork,omitempty" yaml:"vpcSubnetwork,omitempty"`
}
