package types

type Imagebuilder_ContainerRecipeTargetRepository struct {
	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// The service in which this image is registered. Valid values: `ECR`.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
