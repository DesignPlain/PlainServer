package codestarconnections

import types "DesignSphere_Server/src/resource/aws/types"

type Host struct {
	// The name of the host to be created. The name must be unique in the calling AWS account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The endpoint of the infrastructure to be represented by the host after it is created.
	ProviderEndpoint string `json:"providerEndpoint,omitempty" yaml:"providerEndpoint,omitempty"`

	// The name of the external provider where your third-party code repository is configured.
	ProviderType string `json:"providerType,omitempty" yaml:"providerType,omitempty"`

	// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
	VpcConfiguration types.Codestarconnections_HostVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`
}
