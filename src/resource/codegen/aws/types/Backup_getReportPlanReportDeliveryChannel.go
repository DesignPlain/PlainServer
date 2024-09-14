package types

type Backup_getReportPlanReportDeliveryChannel struct {
	// List of the format of your reports: CSV, JSON, or both.
	Formats []string `json:"formats,omitempty" yaml:"formats,omitempty"`

	// Unique name of the S3 bucket that receives your reports.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
