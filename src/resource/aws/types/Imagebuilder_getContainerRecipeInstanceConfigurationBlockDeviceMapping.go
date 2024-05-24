package types

type Imagebuilder_getContainerRecipeInstanceConfigurationBlockDeviceMapping struct {
	// Name of the device. For example, `/dev/sda` or `/dev/xvdb`.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Single list of object with Elastic Block Storage (EBS) block device mapping settings.
	Ebs []Imagebuilder_getContainerRecipeInstanceConfigurationBlockDeviceMappingEb `json:"ebs,omitempty" yaml:"ebs,omitempty"`

	// Whether to remove a mapping from the parent image.
	NoDevice string `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	// Virtual device name. For example, `ephemeral0`. Instance store volumes are numbered starting from 0.
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
