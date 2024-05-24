package types

type Ec2_getLaunchTemplateBlockDeviceMapping struct {
	//
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	Ebs []Ec2_getLaunchTemplateBlockDeviceMappingEb `json:"ebs,omitempty" yaml:"ebs,omitempty"`

	//
	NoDevice string `json:"noDevice,omitempty" yaml:"noDevice,omitempty"`

	//
	VirtualName string `json:"virtualName,omitempty" yaml:"virtualName,omitempty"`
}
