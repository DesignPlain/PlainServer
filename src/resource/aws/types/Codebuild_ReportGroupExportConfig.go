package types

type Codebuild_ReportGroupExportConfig struct {
	// contains information about the S3 bucket where the run of a report is exported. see S3 Destination documented below.
	S3Destination Codebuild_ReportGroupExportConfigS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`

	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
