package networkmanager

import types "DesignSphere_Server/src/resource/aws/types"

type ConnectAttachment struct {
	/*
	   Options block. See options for more information.

	   The following arguments are optional:
	*/
	Options types.Networkmanager_ConnectAttachmentOptions `json:"options,omitempty" yaml:"options,omitempty"`

	// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the attachment between the two connections.
	TransportAttachmentId string `json:"transportAttachmentId,omitempty" yaml:"transportAttachmentId,omitempty"`

	// The ID of a core network where you want to create the attachment.
	CoreNetworkId string `json:"coreNetworkId,omitempty" yaml:"coreNetworkId,omitempty"`

	// The Region where the edge is located.
	EdgeLocation string `json:"edgeLocation,omitempty" yaml:"edgeLocation,omitempty"`
}
