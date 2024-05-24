package dms

type ReplicationSubnetGroup struct {
	// Description for the subnet group.
	ReplicationSubnetGroupDescription string `json:"replicationSubnetGroupDescription,omitempty" yaml:"replicationSubnetGroupDescription,omitempty"`

	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId string `json:"replicationSubnetGroupId,omitempty" yaml:"replicationSubnetGroupId,omitempty"`

	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
