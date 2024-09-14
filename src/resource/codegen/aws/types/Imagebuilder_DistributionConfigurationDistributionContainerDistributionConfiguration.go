package types

type Imagebuilder_DistributionConfigurationDistributionContainerDistributionConfiguration struct {
	// Set of tags that are attached to the container distribution configuration.
	ContainerTags []string `json:"containerTags,omitempty" yaml:"containerTags,omitempty"`

	// Description of the container distribution configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block with the destination repository for the container distribution configuration.
	TargetRepository Imagebuilder_DistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository `json:"targetRepository,omitempty" yaml:"targetRepository,omitempty"`
}
