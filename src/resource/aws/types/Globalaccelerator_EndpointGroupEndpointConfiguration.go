package types

type Globalaccelerator_EndpointGroupEndpointConfiguration struct {
	/*
	   Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint. See the [AWS documentation](https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html) for more details. The default value is `false`.
	   --Note:-- When client IP address preservation is enabled, the Global Accelerator service creates an EC2 Security Group in the VPC named `GlobalAccelerator` that must be deleted (potentially outside of the provider) before the VPC will successfully delete. If this EC2 Security Group is not deleted, the provider will retry the VPC deletion for a few minutes before reporting a `DependencyViolation` error. This cannot be resolved by re-running the provider.
	*/
	ClientIpPreservationEnabled bool `json:"clientIpPreservationEnabled,omitempty" yaml:"clientIpPreservationEnabled,omitempty"`

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`

	// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
