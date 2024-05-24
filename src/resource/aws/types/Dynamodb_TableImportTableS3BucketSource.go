package types

type Dynamodb_TableImportTableS3BucketSource struct {
	// The S3 bucket that is being imported from.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The account number of the S3 bucket that is being imported from.
	BucketOwner string `json:"bucketOwner,omitempty" yaml:"bucketOwner,omitempty"`

	// The key prefix shared by all S3 Objects that are being imported.
	KeyPrefix string `json:"keyPrefix,omitempty" yaml:"keyPrefix,omitempty"`
}
