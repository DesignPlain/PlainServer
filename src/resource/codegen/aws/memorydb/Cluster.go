package memorydb

type Cluster struct {
	// Set of VPC Security Group ID-s to associate with this cluster.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
	SnapshotArns []string `json:"snapshotArns,omitempty" yaml:"snapshotArns,omitempty"`

	// The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to `0`, automatic backups are disabled. Defaults to `0`.
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	// Description for the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to `6379`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The compute and memory capacity of the nodes in the cluster. See AWS documentation on [supported node types](https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.supportedtypes.html) as well as [vertical scaling](https://docs.aws.amazon.com/memorydb/latest/devguide/cluster-vertical-scaling.html).

	   The following arguments are optional:
	*/
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// The name of a snapshot from which to restore data into the new cluster.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`

	// A flag to enable in-transit encryption on the cluster. When set to `false`, the `acl_name` must be `open-access`. Defaults to `true`.
	TlsEnabled bool `json:"tlsEnabled,omitempty" yaml:"tlsEnabled,omitempty"`

	// Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// ARN of the KMS key used to encrypt the cluster at rest.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:23:00-mon:01:30`.
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The number of replicas to apply to each shard, up to a maximum of 5. Defaults to `1` (i.e. 2 nodes per shard).
	NumReplicasPerShard int `json:"numReplicasPerShard,omitempty" yaml:"numReplicasPerShard,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: `05:00-09:00`.
	SnapshotWindow string `json:"snapshotWindow,omitempty" yaml:"snapshotWindow,omitempty"`

	// The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// When set to `true`, the cluster will automatically receive minor engine version upgrades after launch. Defaults to `true`.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotName string `json:"finalSnapshotName,omitempty" yaml:"finalSnapshotName,omitempty"`

	// The number of shards in the cluster. Defaults to `1`.
	NumShards int `json:"numShards,omitempty" yaml:"numShards,omitempty"`

	// The name of the parameter group associated with the cluster.
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// The name of the Access Control List to associate with the cluster.
	AclName string `json:"aclName,omitempty" yaml:"aclName,omitempty"`

	// Enables data tiering. This option is not supported by all instance types. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).
	DataTiering bool `json:"dataTiering,omitempty" yaml:"dataTiering,omitempty"`
}
