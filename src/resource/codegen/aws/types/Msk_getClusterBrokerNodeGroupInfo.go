package types

type Msk_getClusterBrokerNodeGroupInfo struct {
	//
	AzDistribution string `json:"azDistribution,omitempty" yaml:"azDistribution,omitempty"`

	//
	ClientSubnets []string `json:"clientSubnets,omitempty" yaml:"clientSubnets,omitempty"`

	//
	ConnectivityInfos []Msk_getClusterBrokerNodeGroupInfoConnectivityInfo `json:"connectivityInfos,omitempty" yaml:"connectivityInfos,omitempty"`

	//
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	//
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	StorageInfos []Msk_getClusterBrokerNodeGroupInfoStorageInfo `json:"storageInfos,omitempty" yaml:"storageInfos,omitempty"`
}
