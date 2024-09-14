package types

type Cloudsearch_DomainScalingParameters struct {
	// The instance type that you want to preconfigure for your domain. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_ScalingParameters.html) for valid values.
	DesiredInstanceType string `json:"desiredInstanceType,omitempty" yaml:"desiredInstanceType,omitempty"`

	// The number of partitions you want to preconfigure for your domain. Only valid when you select `search.2xlarge` as the instance type.
	DesiredPartitionCount int `json:"desiredPartitionCount,omitempty" yaml:"desiredPartitionCount,omitempty"`

	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount int `json:"desiredReplicationCount,omitempty" yaml:"desiredReplicationCount,omitempty"`
}
