package ec2

type TrafficMirrorTarget struct {
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId string `json:"gatewayLoadBalancerEndpointId,omitempty" yaml:"gatewayLoadBalancerEndpointId,omitempty"`

	// The network interface ID that is associated with the target.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn string `json:"networkLoadBalancerArn,omitempty" yaml:"networkLoadBalancerArn,omitempty"`

	/*
	   Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   --NOTE:-- Either `network_interface_id` or `network_load_balancer_arn` should be specified and both should not be specified together
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description of the traffic mirror session.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
