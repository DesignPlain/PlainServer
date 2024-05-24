package types

type Ec2_getNetworkInterfaceAttachment struct {
	//
	AttachmentId string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`

	//
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	InstanceOwnerId string `json:"instanceOwnerId,omitempty" yaml:"instanceOwnerId,omitempty"`
}
