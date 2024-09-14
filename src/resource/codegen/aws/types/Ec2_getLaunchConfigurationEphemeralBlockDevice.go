package types

type Ec2_getLaunchConfigurationEphemeralBlockDevice struct {
	// Name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Virtual Name of the device.
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
