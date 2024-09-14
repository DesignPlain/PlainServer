package ec2transitgateway

import types "libds/aws/types"

type InstanceConnectEndpoint struct {
	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Ec2transitgateway_InstanceConnectEndpointTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Indicates whether your client's IP address is preserved as the source. Default: `true`.
	PreserveClientIp bool `json:"preserveClientIp,omitempty" yaml:"preserveClientIp,omitempty"`

	// One or more security groups to associate with the endpoint. If you don't specify a security group, the default security group for the VPC will be associated with the endpoint.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the subnet in which to create the EC2 Instance Connect Endpoint.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
