package ec2

type NetworkInterfaceAttachment struct {
	// ENI ID to attach.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// Network interface index (int).
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	// Instance ID to attach.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
