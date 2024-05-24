package types

type Sagemaker_EndpointConfigurationDataCaptureConfigCaptureOption struct {
	// Specifies the data to be captured. Should be one of `Input` or `Output`.
	CaptureMode string `json:"captureMode,omitempty" yaml:"captureMode,omitempty"`
}
