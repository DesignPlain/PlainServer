package types

type Msk_ReplicatorKafkaClusterVpcConfig struct {
	// The AWS security groups to associate with the ENIs used by the replicator. If a security group is not specified, the default security group associated with the VPC is used.
	SecurityGroupsIds []string `json:"securityGroupsIds,omitempty" yaml:"securityGroupsIds,omitempty"`

	// The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets to allow communication between your Kafka Cluster and the replicator.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
