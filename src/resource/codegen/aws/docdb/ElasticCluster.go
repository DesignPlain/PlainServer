package docdb

import types "libds/aws/types"

type ElasticCluster struct {
	// List of VPC security groups to associate with the Elastic DocumentDB Cluster
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Password for the Elastic DocumentDB cluster administrator. Can contain any printable ASCII characters. Must be at least 8 characters
	AdminUserPassword string `json:"adminUserPassword,omitempty" yaml:"adminUserPassword,omitempty"`

	// Name of the Elastic DocumentDB cluster
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Weekly time range during which system maintenance can occur in UTC. Format: `ddd:hh24:mi-ddd:hh24:mi`. If not specified, AWS will choose a random 30-minute window on a random day of the week.
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Number of vCPUs assigned to each elastic cluster shard. Maximum is 64. Allowed values are 2, 4, 8, 16, 32, 64
	ShardCapacity int `json:"shardCapacity,omitempty" yaml:"shardCapacity,omitempty"`

	/*
	   Number of shards assigned to the elastic cluster. Maximum is 32

	   The following arguments are optional:
	*/
	ShardCount int `json:"shardCount,omitempty" yaml:"shardCount,omitempty"`

	// IDs of subnets in which the Elastic DocumentDB Cluster operates.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Docdb_ElasticClusterTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Name of the Elastic DocumentDB cluster administrator
	AdminUserName string `json:"adminUserName,omitempty" yaml:"adminUserName,omitempty"`

	// Authentication type for the Elastic DocumentDB cluster. Valid values are `PLAIN_TEXT` and `SECRET_ARN`
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	// ARN of a KMS key that is used to encrypt the Elastic DocumentDB cluster. If not specified, the default encryption key that KMS creates for your account is used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
