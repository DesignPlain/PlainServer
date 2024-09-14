package dynamodb

type TableExport struct {
	// Time in RFC3339 format from which to export table data. The table export will be a snapshot of the table's state at this point in time. Omitting this value will result in a snapshot from the current time.
	ExportTime string `json:"exportTime,omitempty" yaml:"exportTime,omitempty"`

	// Name of the Amazon S3 bucket to export the snapshot to. See the [AWS Documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/S3DataExport_Requesting.html#S3DataExport_Requesting_Permissions) for information on how configure this S3 bucket.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// ID of the AWS account that owns the bucket the export will be stored in.
	S3BucketOwner string `json:"s3BucketOwner,omitempty" yaml:"s3BucketOwner,omitempty"`

	// Amazon S3 bucket prefix to use as the file name and path of the exported snapshot.
	S3Prefix string `json:"s3Prefix,omitempty" yaml:"s3Prefix,omitempty"`

	// Type of encryption used on the bucket where export data will be stored. Valid values are: `AES256`, `KMS`.
	S3SseAlgorithm string `json:"s3SseAlgorithm,omitempty" yaml:"s3SseAlgorithm,omitempty"`

	// ID of the AWS KMS managed key used to encrypt the S3 bucket where export data will be stored (if applicable).
	S3SseKmsKeyId string `json:"s3SseKmsKeyId,omitempty" yaml:"s3SseKmsKeyId,omitempty"`

	/*
	   ARN associated with the table to export.

	   The following arguments are optional:
	*/
	TableArn string `json:"tableArn,omitempty" yaml:"tableArn,omitempty"`

	// Format for the exported data. Valid values are `DYNAMODB_JSON` or `ION`. See the [AWS Documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/S3DataExport.Output.html#S3DataExport.Output_Data) for more information on these export formats. Default is `DYNAMODB_JSON`.
	ExportFormat string `json:"exportFormat,omitempty" yaml:"exportFormat,omitempty"`
}
