package types

type Ec2_getInstanceEphemeralBlockDevice struct {
	// Physical name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Whether the specified device included in the device mapping was suppressed or not (Boolean).
	NoDevice bool `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	// Virtual device name.
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
