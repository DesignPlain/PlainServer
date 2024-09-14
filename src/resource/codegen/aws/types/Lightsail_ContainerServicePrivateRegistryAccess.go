package types

type Lightsail_ContainerServicePrivateRegistryAccess struct {
	// Describes a request to configure an Amazon Lightsail container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See ECR Image Puller Role below for more details.
	EcrImagePullerRole Lightsail_ContainerServicePrivateRegistryAccessEcrImagePullerRole `json:"ecrImagePullerRole,omitempty" yaml:"ecrImagePullerRole,omitempty"`
}
