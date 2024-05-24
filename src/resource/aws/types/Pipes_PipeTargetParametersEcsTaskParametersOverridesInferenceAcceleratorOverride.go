package types

type Pipes_PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride struct {
	// The Elastic Inference accelerator device name to override for the task. This parameter must match a deviceName specified in the task definition.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// The Elastic Inference accelerator type to use.
	DeviceType string `json:"deviceType,omitempty" yaml:"deviceType,omitempty"`
}
