package types

type Imagebuilder_ContainerRecipeInstanceConfiguration struct {
	// Configuration block(s) with block device mappings for the container recipe. Detailed below.
	BlockDeviceMappings []Imagebuilder_ContainerRecipeInstanceConfigurationBlockDeviceMapping `json:"blockDeviceMappings,omitempty" yaml:"blockDeviceMappings,omitempty"`

	// The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}
