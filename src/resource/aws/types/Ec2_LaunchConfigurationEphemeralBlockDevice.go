package types

type Ec2_LaunchConfigurationEphemeralBlockDevice struct {
	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	//
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
