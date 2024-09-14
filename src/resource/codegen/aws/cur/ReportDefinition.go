package cur

type ReportDefinition struct {
	// Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
	RefreshClosedReports bool `json:"refreshClosedReports,omitempty" yaml:"refreshClosedReports,omitempty"`

	// Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
	ReportName string `json:"reportName,omitempty" yaml:"reportName,omitempty"`

	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// Report path prefix. Limited to 256 characters.
	S3Prefix string `json:"s3Prefix,omitempty" yaml:"s3Prefix,omitempty"`

	// Region of the existing S3 bucket to hold generated reports.
	S3Region string `json:"s3Region,omitempty" yaml:"s3Region,omitempty"`

	// A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
	AdditionalArtifacts []string `json:"additionalArtifacts,omitempty" yaml:"additionalArtifacts,omitempty"`

	// A list of schema elements. Valid values are: `RESOURCES`, `SPLIT_COST_ALLOCATION_DATA`.
	AdditionalSchemaElements []string `json:"additionalSchemaElements,omitempty" yaml:"additionalSchemaElements,omitempty"`

	// Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
	ReportVersioning string `json:"reportVersioning,omitempty" yaml:"reportVersioning,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The frequency on which report data are measured and displayed.  Valid values are: `DAILY`, `HOURLY`, `MONTHLY`.
	TimeUnit string `json:"timeUnit,omitempty" yaml:"timeUnit,omitempty"`
}
