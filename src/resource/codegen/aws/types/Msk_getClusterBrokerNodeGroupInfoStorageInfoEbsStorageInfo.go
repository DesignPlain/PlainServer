package types

type Msk_getClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo struct {
	//
	ProvisionedThroughputs []Msk_getClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput `json:"provisionedThroughputs,omitempty" yaml:"provisionedThroughputs,omitempty"`

	//
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`
}
