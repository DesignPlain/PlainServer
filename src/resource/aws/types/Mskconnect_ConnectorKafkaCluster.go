package types

type Mskconnect_ConnectorKafkaCluster struct {
	// The Apache Kafka cluster to which the connector is connected.
	ApacheKafkaCluster Mskconnect_ConnectorKafkaClusterApacheKafkaCluster `json:"apacheKafkaCluster,omitempty" yaml:"apacheKafkaCluster,omitempty"`
}
