package cloudcontrol

type Resource struct {
	// JSON string matching the CloudFormation resource type schema with desired configuration.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role to assume for operations.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// JSON string of the CloudFormation resource type schema which is used for plan time validation where possible. Automatically fetched if not provided. In large scale environments with multiple resources using the same `type_name`, it is recommended to fetch the schema once via the `aws.cloudformation.CloudFormationType` data source and use this argument to reduce `DescribeType` API operation throttling. This value is marked sensitive only to prevent large plan differences from showing.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	/*
	   CloudFormation resource type name. For example, `AWS::EC2::VPC`.

	   The following arguments are optional:
	*/
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`

	// Identifier of the CloudFormation resource type version.
	TypeVersionId string `json:"typeVersionId,omitempty" yaml:"typeVersionId,omitempty"`
}
