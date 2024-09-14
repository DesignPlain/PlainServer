package types

type Sagemaker_EndpointConfigurationDataCaptureConfig struct {
	// Flag to enable data capture. Defaults to `false`.
	EnableCapture bool `json:"enableCapture,omitempty" yaml:"enableCapture,omitempty"`

	// Portion of data to capture. Should be between 0 and 100.
	InitialSamplingPercentage int `json:"initialSamplingPercentage,omitempty" yaml:"initialSamplingPercentage,omitempty"`

	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt the captured data on Amazon S3.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The content type headers to capture. Fields are documented below.
	CaptureContentTypeHeader Sagemaker_EndpointConfigurationDataCaptureConfigCaptureContentTypeHeader `json:"captureContentTypeHeader,omitempty" yaml:"captureContentTypeHeader,omitempty"`

	// Specifies what data to capture. Fields are documented below.
	CaptureOptions []Sagemaker_EndpointConfigurationDataCaptureConfigCaptureOption `json:"captureOptions,omitempty" yaml:"captureOptions,omitempty"`

	// The URL for S3 location where the captured data is stored.
	DestinationS3Uri string `json:"destinationS3Uri,omitempty" yaml:"destinationS3Uri,omitempty"`
}
