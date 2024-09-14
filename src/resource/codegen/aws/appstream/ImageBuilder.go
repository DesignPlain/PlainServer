package appstream

import types "libds/aws/types"

type ImageBuilder struct {
	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints []types.Appstream_ImageBuilderAccessEndpoint `json:"accessEndpoints,omitempty" yaml:"accessEndpoints,omitempty"`

	// Version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion string `json:"appstreamAgentVersion,omitempty" yaml:"appstreamAgentVersion,omitempty"`

	// Description to display.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Instance type to use when launching the image builder.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   Unique name for the image builder.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig types.Appstream_ImageBuilderVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Human-readable friendly name for the AppStream image builder.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo types.Appstream_ImageBuilderDomainJoinInfo `json:"domainJoinInfo,omitempty" yaml:"domainJoinInfo,omitempty"`

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess bool `json:"enableDefaultInternetAccess,omitempty" yaml:"enableDefaultInternetAccess,omitempty"`

	// ARN of the IAM role to apply to the image builder.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// ARN of the public, private, or shared image to use.
	ImageArn string `json:"imageArn,omitempty" yaml:"imageArn,omitempty"`

	// Name of the image used to create the image builder.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`

	// Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
