package ec2

type SecurityGroupAssociation struct {
	// The ID of the security group to be associated with the VPC endpoint.
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`

	// The ID of the VPC endpoint with which the security group will be associated.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// Whether this association should replace the association with the VPC's default security group that is created when no security groups are specified during VPC endpoint creation. At most 1 association per-VPC endpoint should be configured with `replace_default_association = true`.
	ReplaceDefaultAssociation bool `json:"replaceDefaultAssociation,omitempty" yaml:"replaceDefaultAssociation,omitempty"`
}
