package types

type Msk_ClusterBrokerNodeGroupInfoStorageInfo struct {
	// A block that contains EBS volume information. See below.
	EbsStorageInfo Msk_ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo `json:"ebsStorageInfo,omitempty" yaml:"ebsStorageInfo,omitempty"`
}
