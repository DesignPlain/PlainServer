package mskconnect

import types "DesignSphere_Server/src/resource/aws/types"

type Connector struct {
	// A map of keys to values that represent the configuration for the connector.
	ConnectorConfiguration map[string]string `json:"connectorConfiguration,omitempty" yaml:"connectorConfiguration,omitempty"`

	// Specifies which Apache Kafka cluster to connect to. See below.
	KafkaCluster types.Mskconnect_ConnectorKafkaCluster `json:"kafkaCluster,omitempty" yaml:"kafkaCluster,omitempty"`

	// Details of encryption in transit to the Apache Kafka cluster. See below.
	KafkaClusterEncryptionInTransit types.Mskconnect_ConnectorKafkaClusterEncryptionInTransit `json:"kafkaClusterEncryptionInTransit,omitempty" yaml:"kafkaClusterEncryptionInTransit,omitempty"`

	// Details about log delivery. See below.
	LogDelivery types.Mskconnect_ConnectorLogDelivery `json:"logDelivery,omitempty" yaml:"logDelivery,omitempty"`

	// Specifies which plugins to use for the connector. See below.
	Plugins []types.Mskconnect_ConnectorPlugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs. The types of resources depends on the logic of the connector. For example, a connector that has Amazon S3 as a destination must have permissions that allow it to write to the S3 destination bucket.
	ServiceExecutionRoleArn string `json:"serviceExecutionRoleArn,omitempty" yaml:"serviceExecutionRoleArn,omitempty"`

	// Information about the capacity allocated to the connector. See below.
	Capacity types.Mskconnect_ConnectorCapacity `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// A summary description of the connector.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Details of the client authentication used by the Apache Kafka cluster. See below.
	KafkaClusterClientAuthentication types.Mskconnect_ConnectorKafkaClusterClientAuthentication `json:"kafkaClusterClientAuthentication,omitempty" yaml:"kafkaClusterClientAuthentication,omitempty"`

	// The version of Kafka Connect. It has to be compatible with both the Apache Kafka cluster's version and the plugins.
	KafkaconnectVersion string `json:"kafkaconnectVersion,omitempty" yaml:"kafkaconnectVersion,omitempty"`

	// The name of the connector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies which worker configuration to use with the connector. See below.
	WorkerConfiguration types.Mskconnect_ConnectorWorkerConfiguration `json:"workerConfiguration,omitempty" yaml:"workerConfiguration,omitempty"`
}
