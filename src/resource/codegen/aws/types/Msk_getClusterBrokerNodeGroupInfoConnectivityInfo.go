package types

type Msk_getClusterBrokerNodeGroupInfoConnectivityInfo struct {
	//
	PublicAccesses []Msk_getClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `json:"publicAccesses,omitempty" yaml:"publicAccesses,omitempty"`

	//
	VpcConnectivities []Msk_getClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity `json:"vpcConnectivities,omitempty" yaml:"vpcConnectivities,omitempty"`
}
