package types

type Customerprofiles_DomainMatchingExportingConfigS3Exporting struct {
	// The name of the S3 bucket where Identity Resolution Jobs write result files.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// The S3 key name of the location where Identity Resolution Jobs write result files.
	S3KeyName string `json:"s3KeyName,omitempty" yaml:"s3KeyName,omitempty"`
}
