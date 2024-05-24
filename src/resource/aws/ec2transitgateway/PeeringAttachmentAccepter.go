package ec2transitgateway

type PeeringAttachmentAccepter struct {
	// Key-value tags for the EC2 Transit Gateway Peering Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the EC2 Transit Gateway Peering Attachment to manage.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`
}
