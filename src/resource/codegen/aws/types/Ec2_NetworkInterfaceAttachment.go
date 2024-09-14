package types

type Ec2_NetworkInterfaceAttachment struct {
	//
	AttachmentId string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`

	// Integer to define the devices index.
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	// ID of the instance to attach to.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`
}
