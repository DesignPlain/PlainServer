package types

type Imagebuilder_getDistributionConfigurationDistributionContainerDistributionConfiguration struct {
	// Set of tags that are attached to the container distribution configuration.
	ContainerTags []string `json:"containerTags,omitempty" yaml:"containerTags,omitempty"`

	// Description of the container distribution configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Set of destination repositories for the container distribution configuration.
	TargetRepositories []Imagebuilder_getDistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository `json:"targetRepositories,omitempty" yaml:"targetRepositories,omitempty"`
}
