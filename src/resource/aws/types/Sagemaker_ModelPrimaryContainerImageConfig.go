package types

type Sagemaker_ModelPrimaryContainerImageConfig struct {
	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: `Platform` and `Vpc`.
	RepositoryAccessMode string `json:"repositoryAccessMode,omitempty" yaml:"repositoryAccessMode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	RepositoryAuthConfig Sagemaker_ModelPrimaryContainerImageConfigRepositoryAuthConfig `json:"repositoryAuthConfig,omitempty" yaml:"repositoryAuthConfig,omitempty"`
}
