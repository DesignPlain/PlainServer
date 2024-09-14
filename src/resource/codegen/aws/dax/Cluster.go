package dax

import types "libds/aws/types"

type Cluster struct {
	/*
	   One or more VPC security groups associated
	   with the cluster
	*/
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   List of Availability Zones in which the
	   nodes will be created
	*/
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	/*
	   The type of encryption the
	   cluster's endpoint should support. Valid values are: `NONE` and `TLS`.
	   Default value is `NONE`.
	*/
	ClusterEndpointEncryptionType string `json:"clusterEndpointEncryptionType,omitempty" yaml:"clusterEndpointEncryptionType,omitempty"`

	// Description for the cluster
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A valid Amazon Resource Name (ARN) that identifies
	   an IAM role. At runtime, DAX will assume this role and use the role's
	   permissions to access DynamoDB on your behalf
	*/
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	/*
	   The compute and memory capacity of the nodes. See
	   [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	*/
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	/*
	   The number of nodes in the DAX cluster. A
	   replication factor of 1 will create a single-node cluster, without any read
	   replicas
	*/
	ReplicationFactor int `json:"replicationFactor,omitempty" yaml:"replicationFactor,omitempty"`

	/*
	   Name of the parameter group to associate
	   with this DAX cluster
	*/
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// Encrypt at rest options
	ServerSideEncryption types.Dax_ClusterServerSideEncryption `json:"serverSideEncryption,omitempty" yaml:"serverSideEncryption,omitempty"`

	/*
	   Group identifier. DAX converts this name to
	   lowercase
	*/
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	/*
	   Specifies the weekly time range for when
	   maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	   (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	   `sun:05:00-sun:09:00`
	*/
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	/*
	   An Amazon Resource Name (ARN) of an
	   SNS topic to send DAX notifications to. Example:
	   `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	*/
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" yaml:"notificationTopicArn,omitempty"`

	/*
	   Name of the subnet group to be used for the
	   cluster
	*/
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`
}
