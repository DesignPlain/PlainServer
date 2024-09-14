package types

type Imagebuilder_getDistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository struct {
	// Name of the container repository where the output container image is stored.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// Service in which the image is registered.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
