package types

type Mskconnect_ConnectorKafkaCluster struct {
	// The Apache Kafka cluster to which the connector is connected. See `apache_kafka_cluster` Block for details.
	ApacheKafkaCluster Mskconnect_ConnectorKafkaClusterApacheKafkaCluster `json:"apacheKafkaCluster,omitempty" yaml:"apacheKafkaCluster,omitempty"`
}
