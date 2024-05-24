package types

type Msk_ClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// Access control settings for brokers. See below.
	PublicAccess Msk_ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `json:"publicAccess,omitempty" yaml:"publicAccess,omitempty"`

	// VPC connectivity access control for brokers. See below.
	VpcConnectivity Msk_ClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity `json:"vpcConnectivity,omitempty" yaml:"vpcConnectivity,omitempty"`
}
