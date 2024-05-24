package ec2

type NetworkInsightsPath struct {
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// IP address of the destination resource.
	DestinationIp string `json:"destinationIp,omitempty" yaml:"destinationIp,omitempty"`

	// Destination port to analyze access to.
	DestinationPort int `json:"destinationPort,omitempty" yaml:"destinationPort,omitempty"`

	/*
	   Protocol to use for analysis. Valid options are `tcp` or `udp`.

	   The following arguments are optional:
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// IP address of the source resource.
	SourceIp string `json:"sourceIp,omitempty" yaml:"sourceIp,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
