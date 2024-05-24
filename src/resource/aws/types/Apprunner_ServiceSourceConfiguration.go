package types

type Apprunner_ServiceSourceConfiguration struct {
	// Describes resources needed to authenticate access to some source repositories. See Authentication Configuration below for more details.
	AuthenticationConfiguration Apprunner_ServiceSourceConfigurationAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" yaml:"authenticationConfiguration,omitempty"`

	// Whether continuous integration from the source repository is enabled for the App Runner service. If set to `true`, each repository change (source code commit or new image version) starts a deployment. Defaults to `true`.
	AutoDeploymentsEnabled bool `json:"autoDeploymentsEnabled,omitempty" yaml:"autoDeploymentsEnabled,omitempty"`

	// Description of a source code repository. See Code Repository below for more details.
	CodeRepository Apprunner_ServiceSourceConfigurationCodeRepository `json:"codeRepository,omitempty" yaml:"codeRepository,omitempty"`

	// Description of a source image repository. See Image Repository below for more details.
	ImageRepository Apprunner_ServiceSourceConfigurationImageRepository `json:"imageRepository,omitempty" yaml:"imageRepository,omitempty"`
}
