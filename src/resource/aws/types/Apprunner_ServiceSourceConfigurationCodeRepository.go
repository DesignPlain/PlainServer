package types

type Apprunner_ServiceSourceConfigurationCodeRepository struct {
	// Location of the repository that contains the source code.
	RepositoryUrl string `json:"repositoryUrl,omitempty" yaml:"repositoryUrl,omitempty"`

	// Version that should be used within the source code repository. See Source Code Version below for more details.
	SourceCodeVersion Apprunner_ServiceSourceConfigurationCodeRepositorySourceCodeVersion `json:"sourceCodeVersion,omitempty" yaml:"sourceCodeVersion,omitempty"`

	// The path of the directory that stores source code and configuration files. The build and start commands also execute from here. The path is absolute from root and, if not specified, defaults to the repository root.
	SourceDirectory string `json:"sourceDirectory,omitempty" yaml:"sourceDirectory,omitempty"`

	// Configuration for building and running the service from a source code repository. See Code Configuration below for more details.
	CodeConfiguration Apprunner_ServiceSourceConfigurationCodeRepositoryCodeConfiguration `json:"codeConfiguration,omitempty" yaml:"codeConfiguration,omitempty"`
}
