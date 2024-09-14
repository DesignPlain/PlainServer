package ec2

import types "libds/aws/types"

type VpcEndpoint struct {
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`. Interface type endpoints cannot function without being assigned to a subnet.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType string `json:"vpcEndpointType,omitempty" yaml:"vpcEndpointType,omitempty"`

	// The DNS options for the endpoint. See dns_options below.
	DnsOptions types.Ec2_VpcEndpointDnsOptions `json:"dnsOptions,omitempty" yaml:"dnsOptions,omitempty"`

	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	/*
	   Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`. Most users will want this enabled to allow services within the VPC to automatically use the endpoint.
	   Defaults to `false`.
	*/
	PrivateDnsEnabled bool `json:"privateDnsEnabled,omitempty" yaml:"privateDnsEnabled,omitempty"`

	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `json:"routeTableIds,omitempty" yaml:"routeTableIds,omitempty"`

	/*
	   The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	   If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	*/
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the VPC in which the endpoint will be used.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept bool `json:"autoAccept,omitempty" yaml:"autoAccept,omitempty"`

	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	// Subnet configuration for the endpoint, used to select specific IPv4 and/or IPv6 addresses to the endpoint. See subnet_configuration below.
	SubnetConfigurations []types.Ec2_VpcEndpointSubnetConfiguration `json:"subnetConfigurations,omitempty" yaml:"subnetConfigurations,omitempty"`
}
