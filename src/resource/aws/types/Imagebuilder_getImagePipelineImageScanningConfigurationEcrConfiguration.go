package types

type Imagebuilder_getImagePipelineImageScanningConfigurationEcrConfiguration struct {
	// Tags that are added to the output containers that are scanned
	ContainerTags []string `json:"containerTags,omitempty" yaml:"containerTags,omitempty"`

	// The name of the container repository that Amazon Inspector scans
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
