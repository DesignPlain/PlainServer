package types

type Imagebuilder_ImagePipelineImageScanningConfigurationEcrConfiguration struct {
	//
	ContainerTags []string `json:"containerTags,omitempty" yaml:"containerTags,omitempty"`

	// The name of the repository to scan
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
