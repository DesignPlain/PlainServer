package sagemaker

import types "libds/aws/types"

type Space struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the associated Domain.
	DomainId string `json:"domainId,omitempty" yaml:"domainId,omitempty"`

	// A collection of ownership settings. Required if `space_sharing_settings` is set. See `ownership_settings` Block below.
	OwnershipSettings types.Sagemaker_SpaceOwnershipSettings `json:"ownershipSettings,omitempty" yaml:"ownershipSettings,omitempty"`

	// The name of the space that appears in the SageMaker Studio UI.
	SpaceDisplayName string `json:"spaceDisplayName,omitempty" yaml:"spaceDisplayName,omitempty"`

	// The name of the space.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`

	// A collection of space settings. See `space_settings` Block below.
	SpaceSettings types.Sagemaker_SpaceSpaceSettings `json:"spaceSettings,omitempty" yaml:"spaceSettings,omitempty"`

	// A collection of space sharing settings. Required if `ownership_settings` is set. See `space_sharing_settings` Block below.
	SpaceSharingSettings types.Sagemaker_SpaceSpaceSharingSettings `json:"spaceSharingSettings,omitempty" yaml:"spaceSharingSettings,omitempty"`
}
