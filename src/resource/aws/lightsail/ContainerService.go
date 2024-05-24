package lightsail

import types "DesignSphere_Server/src/resource/aws/types"

type ContainerService struct {
	/*
	   The name for the container service. Names must be of length 1 to 63, and be
	   unique within each AWS Region in your Lightsail account.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The power specification for the container service. The power specifies the amount of memory,
	   the number of vCPUs, and the monthly price of each node of the container service.
	   Possible values: `nano`, `micro`, `small`, `medium`, `large`, `xlarge`.
	*/
	Power string `json:"power,omitempty" yaml:"power,omitempty"`

	// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
	PrivateRegistryAccess types.Lightsail_ContainerServicePrivateRegistryAccess `json:"privateRegistryAccess,omitempty" yaml:"privateRegistryAccess,omitempty"`

	/*
	   The public domain names to use with the container service, such as example.com
	   and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
	   specify are used when you create a deployment with a container configured as the public endpoint of your container
	   service. If you don't specify public domain names, then you can use the default domain of the container service.
	   Defined below.
	*/
	PublicDomainNames types.Lightsail_ContainerServicePublicDomainNames `json:"publicDomainNames,omitempty" yaml:"publicDomainNames,omitempty"`

	/*
	   The scale specification for the container service. The scale specifies the allocated compute
	   nodes of the container service.
	*/
	Scale int `json:"scale,omitempty" yaml:"scale,omitempty"`

	/*
	   Map of container service tags. To tag at launch, specify the tags in the Launch Template. If
	   configured with a provider
	   `default_tags` configuration block
	   present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A Boolean value indicating whether the container service is disabled. Defaults to `false`.
	IsDisabled bool `json:"isDisabled,omitempty" yaml:"isDisabled,omitempty"`
}
