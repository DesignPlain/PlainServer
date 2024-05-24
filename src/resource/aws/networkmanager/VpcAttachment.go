package networkmanager

import types "DesignSphere_Server/src/resource/aws/types"

type VpcAttachment struct {
	// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The ARN of the VPC.

	   The following arguments are optional:
	*/
	VpcArn string `json:"vpcArn,omitempty" yaml:"vpcArn,omitempty"`

	// The ID of a core network for the VPC attachment.
	CoreNetworkId string `json:"coreNetworkId,omitempty" yaml:"coreNetworkId,omitempty"`

	// Options for the VPC attachment.
	Options types.Networkmanager_VpcAttachmentOptions `json:"options,omitempty" yaml:"options,omitempty"`

	// The subnet ARN of the VPC attachment.
	SubnetArns []string `json:"subnetArns,omitempty" yaml:"subnetArns,omitempty"`
}
