package msk

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring types.Msk_ClusterOpenMonitoring `json:"openMonitoring,omitempty" yaml:"openMonitoring,omitempty"`

	// Controls storage mode for supported storage tiers. Valid values are: `LOCAL` or `TIERED`.
	StorageMode string `json:"storageMode,omitempty" yaml:"storageMode,omitempty"`

	// Name of the MSK cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Configuration block for specifying encryption. See below.
	EncryptionInfo types.Msk_ClusterEncryptionInfo `json:"encryptionInfo,omitempty" yaml:"encryptionInfo,omitempty"`

	// Specify the desired Kafka software version.
	KafkaVersion string `json:"kafkaVersion,omitempty" yaml:"kafkaVersion,omitempty"`

	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo types.Msk_ClusterLoggingInfo `json:"loggingInfo,omitempty" yaml:"loggingInfo,omitempty"`

	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes int `json:"numberOfBrokerNodes,omitempty" yaml:"numberOfBrokerNodes,omitempty"`

	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo types.Msk_ClusterBrokerNodeGroupInfo `json:"brokerNodeGroupInfo,omitempty" yaml:"brokerNodeGroupInfo,omitempty"`

	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication types.Msk_ClusterClientAuthentication `json:"clientAuthentication,omitempty" yaml:"clientAuthentication,omitempty"`

	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo types.Msk_ClusterConfigurationInfo `json:"configurationInfo,omitempty" yaml:"configurationInfo,omitempty"`

	// Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring string `json:"enhancedMonitoring,omitempty" yaml:"enhancedMonitoring,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
