package networkmanager

import types "DesignSphere_Server/src/resource/aws/types"

type ConnectPeer struct {
	// The Connect peer BGP options.
	BgpOptions types.Networkmanager_ConnectPeerBgpOptions `json:"bgpOptions,omitempty" yaml:"bgpOptions,omitempty"`

	// The ID of the connection attachment.
	ConnectAttachmentId string `json:"connectAttachmentId,omitempty" yaml:"connectAttachmentId,omitempty"`

	// A Connect peer core network address.
	CoreNetworkAddress string `json:"coreNetworkAddress,omitempty" yaml:"coreNetworkAddress,omitempty"`

	// The inside IP addresses used for BGP peering. Required when the Connect attachment protocol is `GRE`. See `aws.networkmanager.ConnectAttachment` for details.
	InsideCidrBlocks []string `json:"insideCidrBlocks,omitempty" yaml:"insideCidrBlocks,omitempty"`

	/*
	   The Connect peer address.

	   The following arguments are optional:
	*/
	PeerAddress string `json:"peerAddress,omitempty" yaml:"peerAddress,omitempty"`

	// The subnet ARN for the Connect peer. Required when the Connect attachment protocol is `NO_ENCAP`. See `aws.networkmanager.ConnectAttachment` for details.
	SubnetArn string `json:"subnetArn,omitempty" yaml:"subnetArn,omitempty"`

	// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
