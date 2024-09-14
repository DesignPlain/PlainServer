package types

type Apprunner_ServiceSourceConfigurationImageRepository struct {
	// Configuration for running the identified image. See Image Configuration below for more details.
	ImageConfiguration Apprunner_ServiceSourceConfigurationImageRepositoryImageConfiguration `json:"imageConfiguration,omitempty" yaml:"imageConfiguration,omitempty"`

	/*
	   Identifier of an image. For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the
	   image name format, see Pulling an image in the Amazon ECR User Guide.
	*/
	ImageIdentifier string `json:"imageIdentifier,omitempty" yaml:"imageIdentifier,omitempty"`

	// Type of the image repository. This reflects the repository provider and whether the repository is private or public. Valid values: `ECR` , `ECR_PUBLIC`.
	ImageRepositoryType string `json:"imageRepositoryType,omitempty" yaml:"imageRepositoryType,omitempty"`
}
