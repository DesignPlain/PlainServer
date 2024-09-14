package types

type Sagemaker_EndpointConfigurationDataCaptureConfigCaptureOption struct {
	// Specifies the data to be captured. Should be one of `Input`, `Output` or `InputAndOutput`.
	CaptureMode string `json:"captureMode,omitempty" yaml:"captureMode,omitempty"`
}
