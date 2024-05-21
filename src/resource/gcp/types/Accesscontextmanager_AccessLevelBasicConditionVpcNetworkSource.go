package types

type Accesscontextmanager_AccessLevelBasicConditionVpcNetworkSource struct {
	/*
	   Sub networks within a VPC network.
	   Structure is documented below.
	*/
	VpcSubnetwork Accesscontextmanager_AccessLevelBasicConditionVpcNetworkSourceVpcSubnetwork `json:"vpcSubnetwork,omitempty" yaml:"vpcSubnetwork,omitempty"`
}
