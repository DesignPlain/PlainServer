package types

type Dms_ReplicationConfigComputeConfig struct {
	// A list of custom DNS name servers supported for the DMS Serverless replication to access your source or target database.
	DnsNameServers string `json:"dnsNameServers,omitempty" yaml:"dnsNameServers,omitempty"`

	// An Key Management Service (KMS) key Amazon Resource Name (ARN) that is used to encrypt the data during DMS Serverless replication. If you don't specify a value for the KmsKeyId parameter, DMS uses your default encryption key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies the maximum value of the DMS capacity units (DCUs) for which a given DMS Serverless replication can be provisioned. A single DCU is 2GB of RAM, with 2 DCUs as the minimum value allowed. The list of valid DCU values includes 2, 4, 8, 16, 32, 64, 128, 192, 256, and 384.
	MaxCapacityUnits int `json:"maxCapacityUnits,omitempty" yaml:"maxCapacityUnits,omitempty"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
	MultiAz bool `json:"multiAz,omitempty" yaml:"multiAz,omitempty"`

	// Specifies a subnet group identifier to associate with the DMS Serverless replication.
	ReplicationSubnetGroupId string `json:"replicationSubnetGroupId,omitempty" yaml:"replicationSubnetGroupId,omitempty"`

	// The Availability Zone where the DMS Serverless replication using this configuration will run. The default value is a random.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Specifies the minimum value of the DMS capacity units (DCUs) for which a given DMS Serverless replication can be provisioned. The list of valid DCU values includes 2, 4, 8, 16, 32, 64, 128, 192, 256, and 384. If this value isn't set DMS scans the current activity of available source tables to identify an optimum setting for this parameter.
	MinCapacityUnits int `json:"minCapacityUnits,omitempty" yaml:"minCapacityUnits,omitempty"`

	/*
	   The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).

	   - Default: A 30-minute window selected at random from an 8-hour block of time per region, occurring on a random day of the week.
	   - Format: `ddd:hh24:mi-ddd:hh24:mi`
	   - Valid Days: `mon, tue, wed, thu, fri, sat, sun`
	   - Constraints: Minimum 30-minute window.
	*/
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Specifies the virtual private cloud (VPC) security group to use with the DMS Serverless replication. The VPC security group must work with the VPC containing the replication.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`
}
