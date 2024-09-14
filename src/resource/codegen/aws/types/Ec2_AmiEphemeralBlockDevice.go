package types

type Ec2_AmiEphemeralBlockDevice struct {
	// Path at which the device is exposed to created instances.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   Name for the ephemeral device, of the form "ephemeralN" where
	   -N- is a volume number starting from zero.
	*/
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
