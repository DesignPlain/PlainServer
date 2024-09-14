package types

type Ec2_InstanceEphemeralBlockDevice struct {
	// Name of the block device to mount on the instance.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Suppresses the specified device included in the AMI's block device mapping.
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	/*
	   [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g., `ephemeral0`).

	   Each AWS Instance type has a different set of Instance Store block devices available for attachment. AWS [publishes a list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes) of which ephemeral devices are available on each type. The devices are always identified by the `virtual_name` in the format `ephemeral{0..N}`.
	*/
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
