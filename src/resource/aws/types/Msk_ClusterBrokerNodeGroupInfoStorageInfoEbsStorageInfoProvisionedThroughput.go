package types

type Msk_ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput struct {
	// Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second. The minimum value is `250`. The maximum value varies between broker type. You can refer to the valid values for the maximum volume throughput at the following [documentation on throughput bottlenecks](https://docs.aws.amazon.com/msk/latest/developerguide/msk-provision-throughput.html#throughput-bottlenecks)
	VolumeThroughput int `json:"volumeThroughput,omitempty" yaml:"volumeThroughput,omitempty"`

	// Controls whether provisioned throughput is enabled or not. Default value: `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
