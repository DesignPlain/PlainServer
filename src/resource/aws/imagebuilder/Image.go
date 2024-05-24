package imagebuilder

import types "DesignSphere_Server/src/resource/aws/types"

type Image struct {
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn string `json:"distributionConfigurationArn,omitempty" yaml:"distributionConfigurationArn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled bool `json:"enhancedImageMetadataEnabled,omitempty" yaml:"enhancedImageMetadataEnabled,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn string `json:"imageRecipeArn,omitempty" yaml:"imageRecipeArn,omitempty"`

	// Configuration block with image scanning configuration. Detailed below.
	ImageScanningConfiguration types.Imagebuilder_ImageImageScanningConfiguration `json:"imageScanningConfiguration,omitempty" yaml:"imageScanningConfiguration,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration types.Imagebuilder_ImageImageTestsConfiguration `json:"imageTestsConfiguration,omitempty" yaml:"imageTestsConfiguration,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.

	   The following arguments are optional:
	*/
	InfrastructureConfigurationArn string `json:"infrastructureConfigurationArn,omitempty" yaml:"infrastructureConfigurationArn,omitempty"`

	// Key-value map of resource tags for the Image Builder Image. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn string `json:"containerRecipeArn,omitempty" yaml:"containerRecipeArn,omitempty"`
}
