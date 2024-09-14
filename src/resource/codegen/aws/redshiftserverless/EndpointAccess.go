package redshiftserverless

type EndpointAccess struct {
	// An array of VPC subnet IDs to associate with the endpoint.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// An array of security group IDs to associate with the workgroup.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// The name of the workgroup.
	WorkgroupName string `json:"workgroupName,omitempty" yaml:"workgroupName,omitempty"`

	// The name of the endpoint.
	EndpointName string `json:"endpointName,omitempty" yaml:"endpointName,omitempty"`

	// The owner Amazon Web Services account for the Amazon Redshift Serverless workgroup.
	OwnerAccount string `json:"ownerAccount,omitempty" yaml:"ownerAccount,omitempty"`
}
