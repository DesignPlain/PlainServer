package types

type Backup_ReportPlanReportDeliveryChannel struct {
	// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
	Formats []string `json:"formats,omitempty" yaml:"formats,omitempty"`

	// The unique name of the S3 bucket that receives your reports.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// The prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
