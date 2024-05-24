package types

type Codebuild_ProjectLogsConfig struct {
	// Configuration block. Detailed below.
	CloudwatchLogs Codebuild_ProjectLogsConfigCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// Configuration block. Detailed below.
	S3Logs Codebuild_ProjectLogsConfigS3Logs `json:"s3Logs,omitempty" yaml:"s3Logs,omitempty"`
}
