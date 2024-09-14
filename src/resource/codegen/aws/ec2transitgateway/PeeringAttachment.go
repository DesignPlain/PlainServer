package ec2transitgateway

import types "libds/aws/types"

type PeeringAttachment struct {
	// Region of EC2 Transit Gateway to peer with.
	PeerRegion string `json:"peerRegion,omitempty" yaml:"peerRegion,omitempty"`

	// Identifier of EC2 Transit Gateway to peer with.
	PeerTransitGatewayId string `json:"peerTransitGatewayId,omitempty" yaml:"peerTransitGatewayId,omitempty"`

	// Key-value tags for the EC2 Transit Gateway Peering Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// Describes whether dynamic routing is enabled or disabled for the transit gateway peering request. See options below for more details!
	Options types.Ec2transitgateway_PeeringAttachmentOptions `json:"options,omitempty" yaml:"options,omitempty"`

	// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
	PeerAccountId string `json:"peerAccountId,omitempty" yaml:"peerAccountId,omitempty"`
}
