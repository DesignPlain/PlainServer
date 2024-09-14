package msk

import types "libds/aws/types"

type Replicator struct {
	// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
	ReplicationInfoList types.Msk_ReplicatorReplicationInfoList `json:"replicationInfoList,omitempty" yaml:"replicationInfoList,omitempty"`

	// The name of the replicator.
	ReplicatorName string `json:"replicatorName,omitempty" yaml:"replicatorName,omitempty"`

	// The ARN of the IAM role used by the replicator to access resources in the customer's account (e.g source and target clusters).
	ServiceExecutionRoleArn string `json:"serviceExecutionRoleArn,omitempty" yaml:"serviceExecutionRoleArn,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A summary description of the replicator.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A list of Kafka clusters which are targets of the replicator.
	KafkaClusters []types.Msk_ReplicatorKafkaCluster `json:"kafkaClusters,omitempty" yaml:"kafkaClusters,omitempty"`
}
