package types

type Ec2_LaunchTemplateBlockDeviceMapping struct {
	/*
	   The [Instance Store Device
	   Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
	   (e.g., `"ephemeral0"`).
	*/
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`

	// The name of the device to mount.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Configure EBS volume properties.
	Ebs Ec2_LaunchTemplateBlockDeviceMappingEbs `json:"ebs,omitempty" yaml:"ebs,omitempty"`

	// Suppresses the specified device included in the AMI's block device mapping.
	NoDevice string `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`
}
