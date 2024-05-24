package types

type Datasync_TaskTaskReportConfig struct {
	// Configuration block containing the configuration for the Amazon S3 bucket where DataSync uploads your task report. See `s3_destination` below.
	S3Destination Datasync_TaskTaskReportConfigS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`

	// Specifies whether your task report includes the new version of each object transferred into an S3 bucket. This only applies if you enable versioning on your bucket. Keep in mind that setting this to INCLUDE can increase the duration of your task execution. Valid values: `INCLUDE` and `NONE`.
	S3ObjectVersioning string `json:"s3ObjectVersioning,omitempty" yaml:"s3ObjectVersioning,omitempty"`

	// Specifies the type of task report you'd like. Valid values: `SUMMARY_ONLY` and `STANDARD`.
	OutputType string `json:"outputType,omitempty" yaml:"outputType,omitempty"`

	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
	ReportLevel string `json:"reportLevel,omitempty" yaml:"reportLevel,omitempty"`

	// Configuration block containing the configuration of the reporting level for aspects of your task report. See `report_overrides` below.
	ReportOverrides Datasync_TaskTaskReportConfigReportOverrides `json:"reportOverrides,omitempty" yaml:"reportOverrides,omitempty"`
}
