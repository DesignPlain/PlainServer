package imagebuilder

import types "libds/aws/types"

type ImagePipeline struct {
	/*
	   Name of the image pipeline.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Description of the image pipeline.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block with image scanning configuration. Detailed below.
	ImageScanningConfiguration types.Imagebuilder_ImagePipelineImageScanningConfiguration `json:"imageScanningConfiguration,omitempty" yaml:"imageScanningConfiguration,omitempty"`

	// Configuration block with schedule settings. Detailed below.
	Schedule types.Imagebuilder_ImagePipelineSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Key-value map of resource tags for the image pipeline. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration types.Imagebuilder_ImagePipelineImageTestsConfiguration `json:"imageTestsConfiguration,omitempty" yaml:"imageTestsConfiguration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn string `json:"infrastructureConfigurationArn,omitempty" yaml:"infrastructureConfigurationArn,omitempty"`

	// Configuration block with the workflow configuration. Detailed below.
	Workflows []types.Imagebuilder_ImagePipelineWorkflow `json:"workflows,omitempty" yaml:"workflows,omitempty"`

	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn string `json:"containerRecipeArn,omitempty" yaml:"containerRecipeArn,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn string `json:"distributionConfigurationArn,omitempty" yaml:"distributionConfigurationArn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled bool `json:"enhancedImageMetadataEnabled,omitempty" yaml:"enhancedImageMetadataEnabled,omitempty"`

	// Amazon Resource Name (ARN) of the service-linked role to be used by Image Builder to [execute workflows](https://docs.aws.amazon.com/imagebuilder/latest/userguide/manage-image-workflows.html).
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn string `json:"imageRecipeArn,omitempty" yaml:"imageRecipeArn,omitempty"`
}
