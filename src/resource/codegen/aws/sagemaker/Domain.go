package sagemaker

import types "libds/aws/types"

type Domain struct {
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement string `json:"appSecurityGroupManagement,omitempty" yaml:"appSecurityGroupManagement,omitempty"`

	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode string `json:"authMode,omitempty" yaml:"authMode,omitempty"`

	// The domain settings. See `domain_settings` Block below.
	DomainSettings types.Sagemaker_DomainDomainSettings `json:"domainSettings,omitempty" yaml:"domainSettings,omitempty"`

	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType string `json:"appNetworkAccessType,omitempty" yaml:"appNetworkAccessType,omitempty"`

	// The default space settings. See `default_space_settings` Block below.
	DefaultSpaceSettings types.Sagemaker_DomainDefaultSpaceSettings `json:"defaultSpaceSettings,omitempty" yaml:"defaultSpaceSettings,omitempty"`

	// The default user settings. See `default_user_settings` Block below.
	DefaultUserSettings types.Sagemaker_DomainDefaultUserSettings `json:"defaultUserSettings,omitempty" yaml:"defaultUserSettings,omitempty"`

	// The domain name.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See `retention_policy` Block below.
	RetentionPolicy types.Sagemaker_DomainRetentionPolicy `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.

	   The following arguments are optional:
	*/
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
