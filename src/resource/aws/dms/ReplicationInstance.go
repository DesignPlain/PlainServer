package dms

type ReplicationInstance struct {
	/*
	   The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).

	   - Default: A 30-minute window selected at random from an 8-hour block of time per region, occurring on a random day of the week.
	   - Format: `ddd:hh24:mi-ddd:hh24:mi`
	   - Valid Days: `mon, tue, wed, thu, fri, sat, sun`
	   - Constraints: Minimum 30-minute window.
	*/
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// The compute and memory capacity of the replication instance as specified by the replication instance class. See [AWS DMS User Guide](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.Types.html) for available instance sizes and advice on which one to choose.
	ReplicationInstanceClass string `json:"replicationInstanceClass,omitempty" yaml:"replicationInstanceClass,omitempty"`

	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId string `json:"replicationSubnetGroupId,omitempty" yaml:"replicationSubnetGroupId,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage int `json:"allocatedStorage,omitempty" yaml:"allocatedStorage,omitempty"`

	// The type of IP address protocol used by a replication instance. Valid values: `IPV4`, `DUAL`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" yaml:"allowMajorVersionUpgrade,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// The engine version number of the replication instance.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
	MultiAz bool `json:"multiAz,omitempty" yaml:"multiAz,omitempty"`

	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	/*
	   The replication instance identifier. This parameter is stored as a lowercase string.

	   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	   - First character must be a letter.
	   - Cannot end with a hyphen
	   - Cannot contain two consecutive hyphens.
	*/
	ReplicationInstanceId string `json:"replicationInstanceId,omitempty" yaml:"replicationInstanceId,omitempty"`
}
