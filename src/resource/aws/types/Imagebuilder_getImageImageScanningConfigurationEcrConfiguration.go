package types

type Imagebuilder_getImageImageScanningConfigurationEcrConfiguration struct {
	// Set of tags for Image Builder to apply to the output container image that that Amazon Inspector scans.
	ContainerTags []string `json:"containerTags,omitempty" yaml:"containerTags,omitempty"`

	// The name of the container repository that Amazon Inspector scans to identify findings for your container images.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
