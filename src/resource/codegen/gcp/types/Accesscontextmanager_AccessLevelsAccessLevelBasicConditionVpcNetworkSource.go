package types

type Accesscontextmanager_AccessLevelsAccessLevelBasicConditionVpcNetworkSource struct {
	/*
	   Sub networks within a VPC network.
	   Structure is documented below.
	*/
	VpcSubnetwork Accesscontextmanager_AccessLevelsAccessLevelBasicConditionVpcNetworkSourceVpcSubnetwork `json:"vpcSubnetwork,omitempty" yaml:"vpcSubnetwork,omitempty"`
}
