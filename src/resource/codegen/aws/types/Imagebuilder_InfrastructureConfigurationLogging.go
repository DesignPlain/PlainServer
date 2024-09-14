package types

type Imagebuilder_InfrastructureConfigurationLogging struct {
	// Configuration block with S3 logging settings. Detailed below.
	S3Logs Imagebuilder_InfrastructureConfigurationLoggingS3Logs `json:"s3Logs,omitempty" yaml:"s3Logs,omitempty"`
}
