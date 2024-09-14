package types

type Imagebuilder_DistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository struct {
	// The service in which this image is registered. Valid values: `ECR`.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
