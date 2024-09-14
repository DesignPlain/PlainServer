package redshiftserverless

import types "libds/aws/types"

type Workgroup struct {
	// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
	EnhancedVpcRouting bool `json:"enhancedVpcRouting,omitempty" yaml:"enhancedVpcRouting,omitempty"`

	// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries, specified in Redshift Processing Units (RPUs).
	MaxCapacity int `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// The port number on which the cluster accepts incoming connections.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity int `json:"baseCapacity,omitempty" yaml:"baseCapacity,omitempty"`

	// An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
	ConfigParameters []types.Redshiftserverless_WorkgroupConfigParameter `json:"configParameters,omitempty" yaml:"configParameters,omitempty"`

	/*
	   The name of the workgroup.

	   The following arguments are optional:
	*/
	WorkgroupName string `json:"workgroupName,omitempty" yaml:"workgroupName,omitempty"`

	// The name of the namespace.
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`

	// A value that specifies whether the workgroup can be accessed from a public network.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`
}
