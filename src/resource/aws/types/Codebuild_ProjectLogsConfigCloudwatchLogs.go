package types

type Codebuild_ProjectLogsConfigCloudwatchLogs struct {
	// Group name of the logs in CloudWatch Logs.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// Current status of logs in S3 for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Prefix of the log stream name of the logs in CloudWatch Logs.
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`
}
