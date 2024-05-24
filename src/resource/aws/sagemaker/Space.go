package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Space struct {
	// The ID of the associated Domain.
	DomainId string `json:"domainId,omitempty" yaml:"domainId,omitempty"`

	// A collection of ownership settings. See Ownership Settings below.
	OwnershipSettings types.Sagemaker_SpaceOwnershipSettings `json:"ownershipSettings,omitempty" yaml:"ownershipSettings,omitempty"`

	// The name of the space that appears in the SageMaker Studio UI.
	SpaceDisplayName string `json:"spaceDisplayName,omitempty" yaml:"spaceDisplayName,omitempty"`

	// The name of the space.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`

	// A collection of space settings. See Space Settings below.
	SpaceSettings types.Sagemaker_SpaceSpaceSettings `json:"spaceSettings,omitempty" yaml:"spaceSettings,omitempty"`

	// A collection of space sharing settings. See Space Sharing Settings below.
	SpaceSharingSettings types.Sagemaker_SpaceSpaceSharingSettings `json:"spaceSharingSettings,omitempty" yaml:"spaceSharingSettings,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
