package datasync

type Agent struct {
	// Key-value pairs of resource tags to assign to the DataSync Agent. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the VPC (virtual private cloud) endpoint that the agent has access to.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// DataSync Agent activation key during resource creation. Conflicts with `ip_address`. If an `ip_address` is provided instead, the provider will retrieve the `activation_key` as part of the resource creation.
	ActivationKey string `json:"activationKey,omitempty" yaml:"activationKey,omitempty"`

	// DataSync Agent IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. DataSync Agent must be accessible on port 80 from where the provider is running.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// Name of the DataSync Agent.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The IP address of the VPC endpoint the agent should connect to when retrieving an activation key during resource creation. Conflicts with `activation_key`.
	PrivateLinkEndpoint string `json:"privateLinkEndpoint,omitempty" yaml:"privateLinkEndpoint,omitempty"`

	// The ARNs of the security groups used to protect your data transfer task subnets.
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	// The Amazon Resource Names (ARNs) of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	SubnetArns []string `json:"subnetArns,omitempty" yaml:"subnetArns,omitempty"`
}
