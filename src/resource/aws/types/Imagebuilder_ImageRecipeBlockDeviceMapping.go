package types

type Imagebuilder_ImageRecipeBlockDeviceMapping struct {
	// Set to `true` to remove a mapping from the parent image.
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	// Virtual device name. For example, `ephemeral0`. Instance store volumes are numbered starting from 0.
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`

	// Name of the device. For example, `/dev/sda` or `/dev/xvdb`.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	Ebs Imagebuilder_ImageRecipeBlockDeviceMappingEbs `json:"ebs,omitempty" yaml:"ebs,omitempty"`
}
