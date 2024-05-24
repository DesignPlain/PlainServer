package imagebuilder

import types "DesignSphere_Server/src/resource/aws/types"

type ImagePipeline struct {
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled bool `json:"enhancedImageMetadataEnabled,omitempty" yaml:"enhancedImageMetadataEnabled,omitempty"`

	// Configuration block with image scanning configuration. Detailed below.
	ImageScanningConfiguration types.Imagebuilder_ImagePipelineImageScanningConfiguration `json:"imageScanningConfiguration,omitempty" yaml:"imageScanningConfiguration,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration types.Imagebuilder_ImagePipelineImageTestsConfiguration `json:"imageTestsConfiguration,omitempty" yaml:"imageTestsConfiguration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn string `json:"infrastructureConfigurationArn,omitempty" yaml:"infrastructureConfigurationArn,omitempty"`

	/*
	   Name of the image pipeline.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block with schedule settings. Detailed below.
	Schedule types.Imagebuilder_ImagePipelineSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn string `json:"containerRecipeArn,omitempty" yaml:"containerRecipeArn,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn string `json:"distributionConfigurationArn,omitempty" yaml:"distributionConfigurationArn,omitempty"`

	// Key-value map of resource tags for the image pipeline. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Description of the image pipeline.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn string `json:"imageRecipeArn,omitempty" yaml:"imageRecipeArn,omitempty"`
}
