package imagebuilder

import types "libds/aws/types"

type ImageRecipe struct {
	// Configuration block(s) with block device mappings for the image recipe. Detailed below.
	BlockDeviceMappings []types.Imagebuilder_ImageRecipeBlockDeviceMapping `json:"blockDeviceMappings,omitempty" yaml:"blockDeviceMappings,omitempty"`

	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components []types.Imagebuilder_ImageRecipeComponent `json:"components,omitempty" yaml:"components,omitempty"`

	// The image recipe uses this image as a base from which to build your customized image. The value can be the base image ARN or an AMI ID.
	ParentImage string `json:"parentImage,omitempty" yaml:"parentImage,omitempty"`

	// Configuration block for the Systems Manager Agent installed by default by Image Builder. Detailed below.
	SystemsManagerAgent types.Imagebuilder_ImageRecipeSystemsManagerAgent `json:"systemsManagerAgent,omitempty" yaml:"systemsManagerAgent,omitempty"`

	/*
	   The semantic version of the image recipe, which specifies the version in the following format, with numeric values in each position to indicate a specific version: major.minor.patch. For example: 1.0.0.

	   The following attributes are optional:
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Description of the image recipe.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the image recipe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags for the image recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 string `json:"userDataBase64,omitempty" yaml:"userDataBase64,omitempty"`

	// The working directory to be used during build and test workflows.
	WorkingDirectory string `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`
}
