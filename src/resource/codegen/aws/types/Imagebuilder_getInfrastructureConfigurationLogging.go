package types

type Imagebuilder_getInfrastructureConfigurationLogging struct {
	// Nested list of S3 logs settings.
	S3Logs []Imagebuilder_getInfrastructureConfigurationLoggingS3Log `json:"s3Logs,omitempty" yaml:"s3Logs,omitempty"`
}
