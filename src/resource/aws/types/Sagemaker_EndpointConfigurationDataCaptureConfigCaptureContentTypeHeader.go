package types

type Sagemaker_EndpointConfigurationDataCaptureConfigCaptureContentTypeHeader struct {
	// The CSV content type headers to capture.
	CsvContentTypes []string `json:"csvContentTypes,omitempty" yaml:"csvContentTypes,omitempty"`

	// The JSON content type headers to capture.
	JsonContentTypes []string `json:"jsonContentTypes,omitempty" yaml:"jsonContentTypes,omitempty"`
}
