package types

type Memorydb_SnapshotClusterConfiguration struct {
	// Name of the snapshot. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Number of shards in the cluster.
	NumShards int `json:"numShards,omitempty" yaml:"numShards,omitempty"`

	// Name of the parameter group associated with the cluster.
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// Port number on which the cluster accepts connections.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Number of days for which MemoryDB retains automatic snapshots before deleting them.
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`

	// The VPC in which the cluster exists.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The weekly time range during which maintenance on the cluster is performed.
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// Version number of the Redis engine used by the cluster.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Compute and memory capacity of the nodes in the cluster.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of the shard.
	SnapshotWindow string `json:"snapshotWindow,omitempty" yaml:"snapshotWindow,omitempty"`

	// Name of the subnet group used by the cluster.
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	// Description for the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
