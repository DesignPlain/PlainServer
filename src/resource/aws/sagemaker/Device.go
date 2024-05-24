package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Device struct {
	// The device to register with SageMaker Edge Manager. See Device details below.
	Device types.Sagemaker_DeviceDevice `json:"device,omitempty" yaml:"device,omitempty"`

	// The name of the Device Fleet.
	DeviceFleetName string `json:"deviceFleetName,omitempty" yaml:"deviceFleetName,omitempty"`
}
