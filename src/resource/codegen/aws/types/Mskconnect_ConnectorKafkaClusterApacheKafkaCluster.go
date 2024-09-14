package types

type Mskconnect_ConnectorKafkaClusterApacheKafkaCluster struct {
	// The bootstrap servers of the cluster.
	BootstrapServers string `json:"bootstrapServers,omitempty" yaml:"bootstrapServers,omitempty"`

	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster. See `vpc` Block for details.
	Vpc Mskconnect_ConnectorKafkaClusterApacheKafkaClusterVpc `json:"vpc,omitempty" yaml:"vpc,omitempty"`
}
