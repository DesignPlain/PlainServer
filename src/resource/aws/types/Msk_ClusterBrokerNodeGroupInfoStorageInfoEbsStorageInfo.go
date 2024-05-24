package types

type Msk_ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo struct {
	// A block that contains EBS volume provisioned throughput information. To provision storage throughput, you must choose broker type kafka.m5.4xlarge or larger. See below.
	ProvisionedThroughput Msk_ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`

	// The size in GiB of the EBS volume for the data drive on each broker node. Minimum value of `1` and maximum value of `16384`.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`
}
