package types

type Ec2_getAmiBlockDeviceMapping struct {
	// Virtual device name (for instance stores).
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`

	// Physical name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Map containing EBS information, if the device is EBS based. Unlike most object attributes, these are accessed directly (e.g., `ebs.volume_size` or `ebs["volume_size"]`) rather than accessed through the first element of a list (e.g., `ebs[0].volume_size`).
	Ebs map[string]string `json:"ebs,omitempty" yaml:"ebs,omitempty"`

	// Suppresses the specified device included in the block device mapping of the AMI.
	NoDevice string `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`
}
