package types

type Msk_ReplicatorKafkaCluster struct {
	// Details of an Amazon MSK cluster.
	AmazonMskCluster Msk_ReplicatorKafkaClusterAmazonMskCluster `json:"amazonMskCluster,omitempty" yaml:"amazonMskCluster,omitempty"`

	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	VpcConfig Msk_ReplicatorKafkaClusterVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`
}
