package networkmanager

type TransitGatewayRouteTableAttachment struct {
	// The ID of the peer for the attachment.
	PeeringId string `json:"peeringId,omitempty" yaml:"peeringId,omitempty"`

	// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ARN of the transit gateway route table for the attachment.
	TransitGatewayRouteTableArn string `json:"transitGatewayRouteTableArn,omitempty" yaml:"transitGatewayRouteTableArn,omitempty"`
}
