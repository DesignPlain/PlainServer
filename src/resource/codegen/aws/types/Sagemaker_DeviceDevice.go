package types

type Sagemaker_DeviceDevice struct {
	// A description for the device.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the device.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Amazon Web Services Internet of Things (IoT) object name.
	IotThingName string `json:"iotThingName,omitempty" yaml:"iotThingName,omitempty"`
}
