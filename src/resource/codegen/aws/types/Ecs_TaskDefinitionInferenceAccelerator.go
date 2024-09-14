package types

type Ecs_TaskDefinitionInferenceAccelerator struct {
	// Elastic Inference accelerator device name. The deviceName must also be referenced in a container definition as a ResourceRequirement.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Elastic Inference accelerator type to use.
	DeviceType string `json:"deviceType,omitempty" yaml:"deviceType,omitempty"`
}
