package quicksight

import types "libds/aws/types"

type VpcConnection struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// A list of IP addresses of DNS resolver endpoints for the VPC connection.
	DnsResolvers []string `json:"dnsResolvers,omitempty" yaml:"dnsResolvers,omitempty"`

	// The display name for the VPC connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A list of subnet IDs for the VPC connection.

	   The following arguments are optional:
	*/
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The IAM role to associate with the VPC connection.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A list of security group IDs for the VPC connection.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	//
	Timeouts types.Quicksight_VpcConnectionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The ID of the VPC connection.
	VpcConnectionId string `json:"vpcConnectionId,omitempty" yaml:"vpcConnectionId,omitempty"`
}
