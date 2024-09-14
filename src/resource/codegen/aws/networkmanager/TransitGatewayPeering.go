package networkmanager

type TransitGatewayPeering struct {
	// The ID of a core network.
	CoreNetworkId string `json:"coreNetworkId,omitempty" yaml:"coreNetworkId,omitempty"`

	// Key-value tags for the peering. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn string `json:"transitGatewayArn,omitempty" yaml:"transitGatewayArn,omitempty"`
}
