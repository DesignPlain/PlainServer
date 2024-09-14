package types

type Msk_ReplicatorReplicationInfoList struct {
	// Configuration relating to consumer group replication.
	ConsumerGroupReplications []Msk_ReplicatorReplicationInfoListConsumerGroupReplication `json:"consumerGroupReplications,omitempty" yaml:"consumerGroupReplications,omitempty"`

	//
	SourceKafkaClusterAlias string `json:"sourceKafkaClusterAlias,omitempty" yaml:"sourceKafkaClusterAlias,omitempty"`

	// The ARN of the source Kafka cluster.
	SourceKafkaClusterArn string `json:"sourceKafkaClusterArn,omitempty" yaml:"sourceKafkaClusterArn,omitempty"`

	// The type of compression to use writing records to target Kafka cluster.
	TargetCompressionType string `json:"targetCompressionType,omitempty" yaml:"targetCompressionType,omitempty"`

	//
	TargetKafkaClusterAlias string `json:"targetKafkaClusterAlias,omitempty" yaml:"targetKafkaClusterAlias,omitempty"`

	// The ARN of the target Kafka cluster.
	TargetKafkaClusterArn string `json:"targetKafkaClusterArn,omitempty" yaml:"targetKafkaClusterArn,omitempty"`

	// Configuration relating to topic replication.
	TopicReplications []Msk_ReplicatorReplicationInfoListTopicReplication `json:"topicReplications,omitempty" yaml:"topicReplications,omitempty"`
}
