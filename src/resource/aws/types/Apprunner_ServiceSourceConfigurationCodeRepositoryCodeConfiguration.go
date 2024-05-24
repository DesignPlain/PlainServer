package types

type Apprunner_ServiceSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// Basic configuration for building and running the App Runner service. Use this parameter to quickly launch an App Runner service without providing an apprunner.yaml file in the source code repository (or ignoring the file if it exists). See Code Configuration Values below for more details.
	CodeConfigurationValues Apprunner_ServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `json:"codeConfigurationValues,omitempty" yaml:"codeConfigurationValues,omitempty"`

	// Source of the App Runner configuration. Valid values: `REPOSITORY`, `API`. Values are interpreted as follows:
	ConfigurationSource string `json:"configurationSource,omitempty" yaml:"configurationSource,omitempty"`
}
