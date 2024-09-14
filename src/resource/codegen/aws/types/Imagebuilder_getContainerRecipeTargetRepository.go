package types

type Imagebuilder_getContainerRecipeTargetRepository struct {
	// Name of the container repository where the output container image is stored. The name is prefixed by the repository location.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// Service in which this image is registered.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
