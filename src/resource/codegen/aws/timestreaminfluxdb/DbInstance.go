package timestreaminfluxdb

import types "libds/aws/types"

type DbInstance struct {
	// Username of the initial admin user created in InfluxDB. Must start with a letter and can't end with a hyphen or contain two consecutive hyphens. This username will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. Along with `bucket`, `organization`, and `password`, this argument will be stored in the secret referred to by the `influx_auth_parameters_secret_arn` attribute.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	/*
	   List of VPC subnet IDs to associate with the DB instance. Provide at least two VPC subnet IDs in different availability zones when deploying with a Multi-AZ standby.

	   The following arguments are optional:
	*/
	VpcSubnetIds []string `json:"vpcSubnetIds,omitempty" yaml:"vpcSubnetIds,omitempty"`

	// Timestream for InfluxDB DB storage type to read and write InfluxDB data. You can choose between 3 different types of provisioned Influx IOPS included storage according to your workloads requirements: Influx IO Included 3000 IOPS, Influx IO Included 12000 IOPS, Influx IO Included 16000 IOPS. Valid options are: `"InfluxIOIncludedT1"`, `"InfluxIOIncludedT2"`, and `"InfluxIOIncludedT1"`. If you use `"InfluxIOIncludedT2" or "InfluxIOIncludedT3", the minimum value for `allocated_storage` is 400.
	DbStorageType string `json:"dbStorageType,omitempty" yaml:"dbStorageType,omitempty"`

	// Configuration for sending InfluxDB engine logs to a specified S3 bucket.
	LogDeliveryConfiguration types.Timestreaminfluxdb_DbInstanceLogDeliveryConfiguration `json:"logDeliveryConfiguration,omitempty" yaml:"logDeliveryConfiguration,omitempty"`

	// Name that uniquely identifies the DB instance when interacting with the Amazon Timestream for InfluxDB API and CLI commands. This name will also be a prefix included in the endpoint. DB instance names must be unique per customer and per region. The argument must start with a letter, cannot contain consecutive hyphens (`-`) and cannot end with a hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amount of storage in GiB (gibibytes). The minimum value is 20, the maximum value is 16384.
	AllocatedStorage int `json:"allocatedStorage,omitempty" yaml:"allocatedStorage,omitempty"`

	// Name of the initial InfluxDB bucket. All InfluxDB data is stored in a bucket. A bucket combines the concept of a database and a retention period (the duration of time that each data point persists). A bucket belongs to an organization. Along with `organization`, `username`, and `password`, this argument will be stored in the secret referred to by the `influx_auth_parameters_secret_arn` attribute.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// ID of the DB parameter group assigned to your DB instance. If added to an existing Timestream for InfluxDB instance or given a new value, will cause an in-place update to the instance. However, if an instance already has a value for `db_parameter_group_identifier`, removing `db_parameter_group_identifier` will cause the instance to be destroyed and recreated.
	DbParameterGroupIdentifier string `json:"dbParameterGroupIdentifier,omitempty" yaml:"dbParameterGroupIdentifier,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Timestream for InfluxDB DB instance type to run InfluxDB on. Valid options are: `"db.influx.medium"`, `"db.influx.large"`, `"db.influx.xlarge"`, `"db.influx.2xlarge"`, `"db.influx.4xlarge"`, `"db.influx.8xlarge"`, `"db.influx.12xlarge"`, and `"db.influx.16xlarge"`.
	DbInstanceType string `json:"dbInstanceType,omitempty" yaml:"dbInstanceType,omitempty"`

	// Name of the initial organization for the initial admin user in InfluxDB. An InfluxDB organization is a workspace for a group of users. Along with `bucket`, `username`, and `password`, this argument will be stored in the secret referred to by the `influx_auth_parameters_secret_arn` attribute.
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// Password of the initial admin user created in InfluxDB. This password will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. Along with `bucket`, `username`, and `organization`, this argument will be stored in the secret referred to by the `influx_auth_parameters_secret_arn` attribute.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// List of VPC security group IDs to associate with the DB instance.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Specifies whether the DB instance will be deployed as a standalone instance or with a Multi-AZ standby for high availability. Valid options are: `"SINGLE_AZ"`, `"WITH_MULTIAZ_STANDBY"`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	// Configures the DB instance with a public IP to facilitate access. Other resources, such as a VPC, a subnet, an internet gateway, and a route table with routes, are also required to enabled public access, in addition to this argument. See "Usage with Public Internet Access Enabled" for an example configuration with all required resources for public internet access.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	//
	Timeouts types.Timestreaminfluxdb_DbInstanceTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
