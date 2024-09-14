package mq

import types "libds/aws/types"

type Broker struct {
	// Authentication strategy used to secure the broker. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
	AuthenticationStrategy string `json:"authenticationStrategy,omitempty" yaml:"authenticationStrategy,omitempty"`

	// Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker. Must be set when `data_replication_mode` is `CRDR`.
	DataReplicationPrimaryBrokerArn string `json:"dataReplicationPrimaryBrokerArn,omitempty" yaml:"dataReplicationPrimaryBrokerArn,omitempty"`

	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	// Version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions. For example, `5.17.6`.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Configuration block for the maintenance window start time. Detailed below.
	MaintenanceWindowStartTime types.Mq_BrokerMaintenanceWindowStartTime `json:"maintenanceWindowStartTime,omitempty" yaml:"maintenanceWindowStartTime,omitempty"`

	// List of security group IDs assigned to the broker.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// List of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires multiple subnets.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Map of tags to assign to the broker. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// Name of the broker.
	BrokerName string `json:"brokerName,omitempty" yaml:"brokerName,omitempty"`

	// Configuration block for broker configuration. Applies to `engine_type` of `ActiveMQ` and `RabbitMQ` only. Detailed below.
	Configuration types.Mq_BrokerConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Deployment mode of the broker. Valid values are `SINGLE_INSTANCE`, `ACTIVE_STANDBY_MULTI_AZ`, and `CLUSTER_MULTI_AZ`. Default is `SINGLE_INSTANCE`.
	DeploymentMode string `json:"deploymentMode,omitempty" yaml:"deploymentMode,omitempty"`

	// Configuration block containing encryption options. Detailed below.
	EncryptionOptions types.Mq_BrokerEncryptionOptions `json:"encryptionOptions,omitempty" yaml:"encryptionOptions,omitempty"`

	// Broker's instance type. For example, `mq.t3.micro`, `mq.m5.large`.
	HostInstanceType string `json:"hostInstanceType,omitempty" yaml:"hostInstanceType,omitempty"`

	// Defines whether this broker is a part of a data replication pair. Valid values are `CRDR` and `NONE`.
	DataReplicationMode string `json:"dataReplicationMode,omitempty" yaml:"dataReplicationMode,omitempty"`

	// Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for `engine_type` `RabbitMQ`. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
	LdapServerMetadata types.Mq_BrokerLdapServerMetadata `json:"ldapServerMetadata,omitempty" yaml:"ldapServerMetadata,omitempty"`

	// Configuration block for the logging configuration of the broker. Detailed below.
	Logs types.Mq_BrokerLogs `json:"logs,omitempty" yaml:"logs,omitempty"`

	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// Storage type of the broker. For `engine_type` `ActiveMQ`, the valid values are `efs` and `ebs`, and the AWS-default is `efs`. For `engine_type` `RabbitMQ`, only `ebs` is supported. When using `ebs`, only the `mq.m5` broker instance type family is supported.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	/*
	   Configuration block for broker users. For `engine_type` of `RabbitMQ`, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.

	   The following arguments are optional:
	*/
	Users []types.Mq_BrokerUser `json:"users,omitempty" yaml:"users,omitempty"`
}
