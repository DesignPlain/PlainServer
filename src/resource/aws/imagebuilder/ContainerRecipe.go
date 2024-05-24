package imagebuilder

import types "DesignSphere_Server/src/resource/aws/types"

type ContainerRecipe struct {
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType string `json:"containerType,omitempty" yaml:"containerType,omitempty"`

	// The description of the container recipe.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the container recipe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The working directory to be used during build and test workflows.
	WorkingDirectory string `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`

	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData string `json:"dockerfileTemplateData,omitempty" yaml:"dockerfileTemplateData,omitempty"`

	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri string `json:"dockerfileTemplateUri,omitempty" yaml:"dockerfileTemplateUri,omitempty"`

	// The base image for the container recipe.
	ParentImage string `json:"parentImage,omitempty" yaml:"parentImage,omitempty"`

	// Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Version of the container recipe.

	   The following attributes are optional:
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components []types.Imagebuilder_ContainerRecipeComponent `json:"components,omitempty" yaml:"components,omitempty"`

	// The KMS key used to encrypt the container image.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration types.Imagebuilder_ContainerRecipeInstanceConfiguration `json:"instanceConfiguration,omitempty" yaml:"instanceConfiguration,omitempty"`

	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride string `json:"platformOverride,omitempty" yaml:"platformOverride,omitempty"`

	// The destination repository for the container image. Detailed below.
	TargetRepository types.Imagebuilder_ContainerRecipeTargetRepository `json:"targetRepository,omitempty" yaml:"targetRepository,omitempty"`
}
