package ec2

type VpcEndpointServiceAllowedPrinciple struct {
	// The ARN of the principal to allow permissions.
	PrincipalArn string `json:"principalArn,omitempty" yaml:"principalArn,omitempty"`

	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId string `json:"vpcEndpointServiceId,omitempty" yaml:"vpcEndpointServiceId,omitempty"`
}
