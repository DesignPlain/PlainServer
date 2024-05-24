package appconfig

type Deployment struct {
	// Application ID. Must be between 4 and 7 characters in length.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Configuration profile ID. Must be between 4 and 7 characters in length.
	ConfigurationProfileId string `json:"configurationProfileId,omitempty" yaml:"configurationProfileId,omitempty"`

	// Configuration version to deploy. Can be at most 1024 characters.
	ConfigurationVersion string `json:"configurationVersion,omitempty" yaml:"configurationVersion,omitempty"`

	// Deployment strategy ID or name of a predefined deployment strategy. See [Predefined Deployment Strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html#appconfig-creating-deployment-strategy-predefined) for more details.
	DeploymentStrategyId string `json:"deploymentStrategyId,omitempty" yaml:"deploymentStrategyId,omitempty"`

	// Description of the deployment. Can be at most 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Environment ID. Must be between 4 and 7 characters in length.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	// The KMS key identifier (key ID, key alias, or key ARN). AppConfig uses this to encrypt the configuration data using a customer managed key.
	KmsKeyIdentifier string `json:"kmsKeyIdentifier,omitempty" yaml:"kmsKeyIdentifier,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
