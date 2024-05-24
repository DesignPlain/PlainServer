package types

type Imagebuilder_getContainerRecipeInstanceConfiguration struct {
	// AMI ID of the base image for container build and test instance.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Set of objects with block device mappings for the instance configuration.
	BlockDeviceMappings []Imagebuilder_getContainerRecipeInstanceConfigurationBlockDeviceMapping `json:"blockDeviceMappings,omitempty" yaml:"blockDeviceMappings,omitempty"`
}
